package notifications

import (
	"net/http"
	"sosmed/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type NotificationHandler struct {
	service INotificationService
	logger  *logrus.Logger
	 upgrader websocket.Upgrader
}


func NewNotificationController(service INotificationService, logger *logrus.Logger) *NotificationHandler {
    return &NotificationHandler{
        service: service,
        logger:  logger,
        upgrader: websocket.Upgrader{
            CheckOrigin: func(r *http.Request) bool {
                return true
            },
        },
    }
}

func (ctrl *NotificationHandler) HandleWebSocket(c *gin.Context) {
    // Ambil token dari query parameter
    token := c.Query("token")
    if token == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
        return
    }

    // Verifikasi token
    claims, err := utils.ValidateToken(token)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    userIDFloat, ok := claims["userId"].(float64)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userId"})
        return
    }
    userID := int(userIDFloat)

    // Upgrade ke WebSocket
    conn, err := ctrl.upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        ctrl.logger.Error("Failed to upgrade to WebSocket: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to WebSocket"})
        return
    }

    // Daftarkan client
    ctrl.service.RegisterClient(userID, conn)
    defer func() {
        ctrl.service.UnregisterClient(userID)
        conn.Close()
    }()

    // Loop untuk menjaga koneksi tetap hidup
    for {
        if _, _, err := conn.ReadMessage(); err != nil {
            ctrl.logger.Info("WebSocket closed for user: ", userID)
            break
        }
    }
}

func (ctrl *NotificationHandler) GetNotifications(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    notifications, err := ctrl.service.GetNotifications(userID.(int))
    if err != nil {
        ctrl.logger.Error("Failed to get notifications: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get notifications"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}