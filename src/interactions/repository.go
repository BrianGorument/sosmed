package interactions

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// userRepository struct
type interactionRepository struct {
	db *gorm.DB
}

// NewUserRepository (Dependency Injection)
func NewInteractionRepository(db *gorm.DB) IInteractionRepository {
	return &interactionRepository{db: db}
}

func (r *interactionRepository) InsertInteraction(input Post_Interactions) (*InteractResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use GORM with context
	if err := r.db.WithContext(ctx).Create(&input).Error; err != nil {
		return nil, err
	}

	return &InteractResponse{
		ID: input.ID ,
		PostID: input.PostID,
		UserID: input.UserID,
		}, nil
}

