package interactions

import (
	"log"
	"sosmed/shared/utils"
)

// userService struct
type interactionsService struct {
	repo IInteractionRepository
}

// NewUserService (Dependency Injection)
func NewInteractionsService(repo IInteractionRepository) IInteractionService {
	return &interactionsService{repo}
}

func (u *interactionsService) InsertOrUpdateInteraction(req InteractRequest, user UserData) (*InteractResponse, error) {
	var resp *InteractResponse
	
	tx, err := u.repo.BeginTransaction()
	if err != nil {
		return resp, err
	}
	defer func() {
		if r := recover(); r != nil {
			u.repo.RollbackTransaction(tx)
			log.Println("Transaction rolled back due to error:", r)
		}
	}()
	
	if req.Comment != "" {
		Post_Interactions := Comments{
			PostID:    req.PostID,
			UserID:    user.UserId,
			Comment:   req.Comment,
			Media:     req.Media,  // Image URL/path
		}
		
		resp , err = u.repo.InsertComment(tx,Post_Interactions)
		if err != nil {
			u.repo.RollbackTransaction(tx)
			return resp, err
		}
	}
	
	if req.Type != "" {
		likes := Likes{
			PostID:    req.PostID,
			UserID:    user.UserId,
			Type:      req.Type,
		}
		resp, err = u.repo.UpdateLikesInteraction(tx, likes)
		if err != nil {
			u.repo.RollbackTransaction(tx)
			return resp, err
		}
	}	
	//commits
	if err := u.repo.CommitTransaction(tx); err != nil {
		u.repo.RollbackTransaction(tx)
		log.Println("Error committing transaction:", err)
		return resp, err
	}

	return resp , nil
}

func (u *interactionsService) DeleteCommentOrMedia(req DeleteCommentRequest, user UserData) (*InteractResponse, error) {
	var resp *InteractResponse
	
	
	tx, err := u.repo.BeginTransaction()
	if err != nil {
		return resp, err
	}
	defer func() {
		if r := recover(); r != nil {
			u.repo.RollbackTransaction(tx)
			log.Println("Transaction rolled back due to error:", r)
		}
	}()	
		deleteComment := Comments{
			ID :       req.ID,
			PostID:    req.PostID,
		}
		
		resp , err = u.repo.DeleteCommentByID(tx,deleteComment)
		if err != nil {
			u.repo.RollbackTransaction(tx)
			return resp, err
		}
	
	//commits
	if err := u.repo.CommitTransaction(tx); err != nil {
		u.repo.RollbackTransaction(tx)
		log.Println("Error committing transaction:", err)
		return resp, err
	}

	return resp , nil
}