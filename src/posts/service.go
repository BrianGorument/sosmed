package posts

import (
	"sosmed/shared/utils"
	"strconv"
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
	
	ImageURL, err := utils.HandleMedia(req.Image)
	if err != nil {
		return postres, err
	}
	stringuserID := strconv.Itoa(int(users.UserId))

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
