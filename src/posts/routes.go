package posts

import (
	"sosmed/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// RegisterRoutes
func RegisterRoutes(router *gin.Engine, db *gorm.DB, log *logrus.Logger) {
	repo := NewPostRepository(db)
	service := NewPostService(repo)
	handler := NewPostingHandler(service, log)

	routersGroup := router.Group("v1")
	{
		postsGroup := routersGroup.Group("posts")

		// usersGroup.GET("/", handler.GetAllUsers)
		postsGroup.POST("/postCreate", utils.JWTAuthMiddleware() , handler.CreatePost)
		postsGroup.GET("/posts/:id", utils.JWTAuthMiddleware(),handler.GetAllPosts)
		postsGroup.POST("/upload/media", utils.JWTAuthMiddleware() , handler.UploadMedia)

	}
}
