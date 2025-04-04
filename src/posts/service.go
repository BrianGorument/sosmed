package users

import (
	"errors"
	"sosmed/shared/utils"
	"strconv"
	"strings"
)

// userService struct
type postService struct {
	repo IPostRepository
}

// NewUserService (Dependency Injection)
func NewPostService(repo IPostRepository) IPostService {
	return &postService{repo}
}

func (s *postService) CreatePosting(req CreatePostRequest , users UserData) (*PostResponse, error) {
	var postres *PostResponse
	
	ImageURL, err := s.handleMedia(req.Image)
	if err != nil {
		return postres, err
	}
	stringuserID := strconv.FormatFloat(users.UserId, 'f', 6, 64)

	post := Post_Content{
		UserID:   stringuserID,
		Title:    req.Title,
		Content:  req.Content,
		Image:    ImageURL,  // Image URL/path
		//Media:    mediaURL,     // Media URL/path (video/image URL or file path)
		LikeCount: 0,
		CategoryId: 0,
	}
	
	postres , err = s.repo.InsertPosting(post)
	if err != nil {
		return postres, err
	}

	return postres , nil
}

func (s *postService) handleMedia(media string) (string, error) {
	if utils.IsValidURL(media) {
		return media, nil
	}

	// Check if the media is a base64 encoded string (assume image or video)
	if strings.HasPrefix(media, "data:image") || strings.HasPrefix(media, "data:video") {
		// Decode the base64 string and save as a file
		decodedMedia, err := utils.DecodeBase64ToFile(media)
		if err != nil {
			return "", errors.New("failed to decode base64 media")
		}
		return decodedMedia, nil
	}

	return "", errors.New("invalid media format")
}