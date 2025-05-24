package notifications

import (
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type INotificationService interface {
	RegisterClient(userID int, conn *websocket.Conn)
	UnregisterClient(userID int)
	SendNotification(userID int,notification Notification) error
	GetNotifications(userID int) ([]Notification, error) 
}

type INotificationRepository interface {
	BeginTransaction() (*gorm.DB, error)               // Untuk memulai transaksi
	CommitTransaction(tx *gorm.DB) error               // Untuk commit transaksi
	RollbackTransaction(tx *gorm.DB) error 
	Save(tx *gorm.DB, notification Notification) error
    GetByUserID(tx *gorm.DB, userID int) ([]Notification, error)
}
