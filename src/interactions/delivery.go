package interactions

import (
	"net/http"
	"sosmed/shared/response"
	"sosmed/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


type InteractionHandler struct {
	service IInteractionService
	logger  *logrus.Logger
}


func NewInteractionHandler(service IInteractionService, logger *logrus.Logger) *InteractionHandler {
	return &InteractionHandler{service, logger}
}

func (h *InteractionHandler) CreateComment(c *gin.Context) {
	var req InteractRequest
	tokenString := c.GetHeader("Authorization")

	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	
	userID, err := utils.ConvertToUint(claims["userId"])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userId type"})
		return
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Invalid request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body request"})
		return
	}
	
	user := UserData{
		UserId:    userID,
	}

	comment, err := h.service.CreateCommentService(req, user)
	if err != nil {
		h.logger.Error("Failed to comment on post:", err)
		resp := response.ErrorStruct{
			Description:        response.DescriptionFailed,
			Message:            err.Error(),
			MessageDescription: "Failed to comment on post",
			Data:               err,
		}
		response.SendErrorResponse(c, http.StatusBadRequest, resp)
		return
	}
	
	succesresp := response.Response{
		ResponseCode:       response.RCSuccess,
		Description:        response.DescriptionSuccess,
		Message:            response.DataSuccess,
		MessageDescription: "Successfully comment on post",
		Data:               comment,
	}
	response.SendResponseSuccess(c, http.StatusOK, succesresp)
	
	
	
}