package interactions

import (
	"sosmed/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// RegisterRoutes
func RegisterRoutes(router *gin.Engine, db *gorm.DB, log *logrus.Logger) {
	repo := NewInteractionRepository(db)
	service := NewInteractionsService(repo)
	handler := NewInteractionHandler(service, log)

	routersGroup := router.Group("v1")
	{
		interactionsGroup := routersGroup.Group("activity")

		// usersGroup.GET("/", handler.GetAllUsers)
		interactionsGroup.POST("/comment", utils.JWTAuthMiddleware() , handler.CreateComment)
		interactionsGroup.POST("/deleteComment", utils.JWTAuthMiddleware() , handler.DeleteComment)

	}
}
