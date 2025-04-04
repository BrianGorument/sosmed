package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// userRepository struct
type postRepository struct {
	db *gorm.DB
}

// NewUserRepository (Dependency Injection)
func NewPostRepository(db *gorm.DB) IPostRepository {
	return &postRepository{db: db}
}


func (r *postRepository) InsertPosting(input Post_Content) (*PostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use GORM with context
	if err := r.db.WithContext(ctx).Create(&input).Error; err != nil {
		return nil, err
	}

	return &PostResponse{Post_ID: input.ID , UserID: input.UserID}, nil
}
