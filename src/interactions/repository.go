package interactions

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// userRepository struct
type interactionRepository struct {
	DB *gorm.DB
}

// NewUserRepository (Dependency Injection)
func NewInteractionRepository(db *gorm.DB) IInteractionRepository {
	return &interactionRepository{DB: db}
}


func (r *interactionRepository) BeginTransaction() (*gorm.DB, error) {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (r *interactionRepository) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *interactionRepository) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *interactionRepository) InsertComment(tx *gorm.DB, input Comments) (*InteractResponse, error) {	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use GORM with context
	if err := tx.WithContext(ctx).Create(&input).Error; err != nil {
		return nil, err
	}

	return &InteractResponse{
		ID: input.ID ,
		PostID: input.PostID,
		UserID: input.UserID,
		}, nil
}

func (r *interactionRepository) UpdateLikesInteraction(tx *gorm.DB, input Likes) (*InteractResponse, error) {
	var table Likes
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := tx.Where("post_id = ? AND user_id = ?", input.PostID, input.UserID).First(&table).Error; err != nil {
		return nil, err
	}
	table.Type = input.Type

	if err := tx.WithContext(ctx).Save(&table).Error; err != nil {
		return nil , err
	}

	return &InteractResponse{
		ID: input.ID ,
		PostID: input.PostID,
		UserID: input.UserID,
		}, nil
}

