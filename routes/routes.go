package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// InitRoutes initializes all routes
func InitRoutes(db *gorm.DB, log *logrus.Logger) *gin.Engine {
	router := gin.Default()

	// Example Route
	router.GET("/ping", func(c *gin.Context) {
		log.Info("Ping endpoint called")
		c.JSON(200, gin.H{"message": "pong"})
	})

	return router
}
