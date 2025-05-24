package notifications

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// userService struct
type notificationsService struct {
	repo INotificationRepository
	logger     *logrus.Logger
    clients    map[int]*websocket.Conn
    clientsMu  sync.Mutex    
}

// NewUserService (Dependency Injection)
func NewNotificationsService(repo INotificationRepository, logger *logrus.Logger , client *websocket.Conn) INotificationService {
	return &notificationsService{
		repo:    repo,
        logger:  logger,
        clients: make(map[int]*websocket.Conn),
	}
	    
}


func (s *notificationsService) RegisterClient(userID int, conn *websocket.Conn) {
    s.clientsMu.Lock()
    s.clients[userID] = conn
    s.clientsMu.Unlock()
}

func (s *notificationsService) UnregisterClient(userID int) {
    s.clientsMu.Lock()
    delete(s.clients, userID)
    s.clientsMu.Unlock()
}

func (s *notificationsService) SendNotification(userID int, notification Notification) error {
    tx, err := s.repo.BeginTransaction()
    if err != nil {
        return err
    }
    defer func() {
        if r := recover(); r != nil {
            s.repo.RollbackTransaction(tx)
            s.logger.Error("Transaction rolled back due to panic: ", r)
        }
    }()

    notification.CreatedAt = time.Now()
    if err := s.repo.Save(tx, notification); err != nil {
        s.repo.RollbackTransaction(tx)
        s.logger.Error("Failed to save notification: ", err)
        return err
    }

    if err := s.repo.CommitTransaction(tx); err != nil {
        s.repo.RollbackTransaction(tx)
        s.logger.Error("Error committing transaction: ", err)
        return err
    }

    s.clientsMu.Lock()
    conn, exists := s.clients[userID]
    s.clientsMu.Unlock()

    if exists {
        msg, err := json.Marshal(notification)
        if err != nil {
            s.logger.Error("Failed to marshal notification: ", err)
            return err
        }
        if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
            s.logger.Error("Failed to send WebSocket message: ", err)
            s.UnregisterClient(userID)
            return err
        }
    }
    return nil
}

func (s *notificationsService) GetNotifications(userID int) ([]Notification, error) {
    tx, err := s.repo.BeginTransaction()
    if err != nil {
        return nil, err
    }
    defer func() {
        if r := recover(); r != nil {
            s.repo.RollbackTransaction(tx)
            s.logger.Error("Transaction rolled back due to panic: ", r)
        }
    }()

    notifications, err := s.repo.GetByUserID(tx, userID)
    if err != nil {
        s.repo.RollbackTransaction(tx)
        s.logger.Error("Failed to get notifications: ", err)
        return nil, err
    }

    if err := s.repo.CommitTransaction(tx); err != nil {
        s.repo.RollbackTransaction(tx)
        s.logger.Error("Error committing transaction: ", err)
        return nil, err
    }

    return notifications, nil
}