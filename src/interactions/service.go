package interactions

import "sosmed/shared/utils"

// userService struct
type interactionsService struct {
	repo IInteractionRepository
}

// NewUserService (Dependency Injection)
func NewInteractionsService(repo IInteractionRepository) IInteractionService {
	return &interactionsService{repo}
}

func (u *interactionsService) CreateCommentService(req InteractRequest, user UserData) (*InteractResponse, error) {
	var resp *InteractResponse
	
	
	if req.Media != "" {
		MediaURL, err := utils.HandleMedia(req.Media)
		if err != nil {
			return resp, err
		}
		req.Media = MediaURL
	}
	
	Post_Interactions := Post_Interactions{
		PostID:    req.PostID,
		UserID:    user.UserId,
		Type:      req.Type,
		Comment:   req.Comment,
		Media:     req.Media,  // Image URL/path
	}

	resp , err := u.repo.InsertInteraction(Post_Interactions)
	if err != nil {
		return resp, err
	}

	return resp , nil
}
