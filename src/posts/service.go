package posts

import (
	"log"
	"sosmed/shared/utils"
	"sosmed/src/interactions"
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
	
	tx, err := s.repo.BeginTransaction()
	if err != nil {
		return postres, err
	}
	defer func() {
		if r := recover(); r != nil {
			s.repo.RollbackTransaction(tx)
			log.Println("Transaction rolled back due to error:", r)
		}
	}()
	
	
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
	
	postres , err = s.repo.InsertPosting(tx,post)
	if err != nil {
		s.repo.RollbackTransaction(tx)
		return postres, err
	}
	
	crtDatalikes := interactions.Likes{
		PostID:    postres.Post_ID,
		UserID:    users.UserId,
		Type: 	   "NULL",
	}
	
	err = s.repo.InsertLikesTable(tx , crtDatalikes)
	if err != nil {
		s.repo.RollbackTransaction(tx)
		return postres, err
	}
	
	if err := s.repo.CommitTransaction(tx); err != nil {
		s.repo.RollbackTransaction(tx)
		log.Println("Error committing transaction:", err)
		return postres, err
	}

	return postres , nil
}

func (s *postService) GetAllPosts(filter GetAllPostsFilterRequest , user UserData) (*GetAllPostsResponse , error) {
	return s.repo.FindAll(filter , user)
}