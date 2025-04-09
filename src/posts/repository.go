package posts

import (
	"context"
	"sosmed/src/interactions"
	"time"

	"gorm.io/gorm"
)

// userRepository struct
type postRepository struct {
	db *gorm.DB
}

// NewUserRepository (Dependency Injection)
func NewPostRepository(DB *gorm.DB) IPostRepository {
	return &postRepository{db: DB}
}

func (r *postRepository) BeginTransaction() (*gorm.DB, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (r *postRepository) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *postRepository) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *postRepository) InsertPosting(tx *gorm.DB, input Post_Content) (*PostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use GORM with context
	if err := tx.WithContext(ctx).Create(&input).Error; err != nil {
		return nil, err
	}

	return &PostResponse{Post_ID: input.ID , UserID: input.UserID}, nil
}


func (r *postRepository) InsertLikesTable(tx *gorm.DB, input interactions.Likes) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use GORM with context
	if err := tx.WithContext(ctx).Create(&input).Error; err != nil {
		return err
	}

	return nil	
}