package users

import (
	"net/http"
	"sosmed/shared/response"
	"sosmed/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


type PostingHandler struct {
	service IPostService
	logger  *logrus.Logger
}


func NewPostingHandler(service IPostService, logger *logrus.Logger) *PostingHandler {
	return &PostingHandler{service, logger}
}

func (h *PostingHandler) CreatePost(c *gin.Context) {
	var req CreatePostRequest
	var users UserData
	
	tokenString := c.GetHeader("Authorization")

	// Verifikasi dan ambil klaim dari token
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Ambil klaim dari token
	userID := claims["userId"].(float64) // Sesuaikan dengan klaim yang ada di token
	userName := claims["userName"].(string)
	userEmail := claims["userEmail"].(string)
	
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Invalid request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body request"})
		return
	}
	
	// Get userID from JWT context
	// userID, exists := c.Get("userId")
	// userName, exists := c.Get("userName")
	// userEmail, exists := c.Get("userEmail")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
	// 	return
	// }
	
	users = UserData{
		UserId:    userID,
		UserEmail: userEmail,
		Username:  userName,
	}

	
	posting, err := h.service.CreatePosting(req , users)
	if err != nil {
		h.logger.Error("Failed to create post:", err)
		resp := response.ErrorStruct{
			Description:        response.DescriptionFailed,
			Message:            err.Error(),
			MessageDescription: "Failed to create post",
			Data:               err,
		}
		response.SendErrorResponse(c, http.StatusBadRequest, resp)
		return
	}
	
	succesresp := response.Response{
		ResponseCode:       response.RCSuccess,
		Description:        response.DescriptionSuccess,
		Message:            response.DataSuccess,
		MessageDescription: "Successfully created post",
		Data:               posting,
	}
	response.SendResponseSuccess(c, http.StatusOK, succesresp)
	
	
	
}