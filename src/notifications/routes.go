package notifications

import (
	"sosmed/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, log *logrus.Logger , client *websocket.Conn) {
	repo := NewNotificationRepository(db)
	service := NewNotificationsService(repo ,log , client)
	handler := NewNotificationController(service, log)
	routersGroup := router.Group("v1")
	{
		notificationsGroup := routersGroup.Group("notifications")

		notificationsGroup.GET("/ws", utils.JWTAuthMiddleware() , handler.GetNotifications)

	}
}
