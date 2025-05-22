package posts

import (
	"net/http"
	"sosmed/shared/response"
	"sosmed/shared/utils"
	"strconv"

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
	
	userID, err := utils.ConvertToUint(claims["userId"])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userId type"})
		return
	}

	userName := claims["userName"].(string)
	userEmail := claims["userEmail"].(string)
	
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Invalid request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body request"})
		return
	}	
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
		Message:            response.SuccessInsert,
		MessageDescription: "Successfully created post",
		Data:               posting,
	}
	response.SendResponseSuccess(c, http.StatusOK, succesresp)
}
func (h *PostingHandler) GetAllPosts(c *gin.Context) {
	var filter GetAllPostsFilterRequest
	filter.Limit , _ = strconv.Atoi(c.DefaultQuery("limit", "10"))
	filter.Page , _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Title = c.Query("title")
	filter.ByUserName = c.Query("userName")
	
	intId , _ := strconv.Atoi(c.Param("id"))
	filter.PostID = uint(intId)
	
	var users UserData
	tokenString := c.GetHeader("Authorization")

	// Verifikasi dan ambil klaim dari token
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

	userName := claims["userName"].(string)
	userEmail := claims["userEmail"].(string)
	
	users = UserData{
		UserId:    userID,
		UserEmail: userEmail,
		Username:  userName,
	}
	
	result, err := h.service.GetAllPosts(filter , users)
	if err != nil {
		h.logger.Error("Failed to get posting data:", err)
		resp := response.ErrorStruct{
			Description:        response.DescriptionFailed,
			Message:            err.Error(),
			MessageDescription: "Failed to get posting data",
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
		Data:               result,
	}
	response.SendResponseSuccess(c, http.StatusOK, succesresp)
}
func (h *PostingHandler) UploadMedia(c *gin.Context) {
    // Ambil file dari form-data
    file, err := c.FormFile("media")
    if err != nil {
        h.logger.Error("Failed to get file: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Media file is required"})
        return
    }
    // Panggil fungsi util untuk upload dan kompresi
    mediaPath, err := utils.UploadAndCompressMedia(file)
    if err != nil {
        h.logger.Error("Failed to upload media: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    h.logger.Info("Media uploaded successfully: ", mediaPath)
    c.JSON(http.StatusOK, gin.H{"message": "Media uploaded successfully", "media_url": mediaPath})
}