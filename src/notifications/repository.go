package notifications

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type notificationRepository struct {
	DB *gorm.DB
}

// NewUserRepository (Dependency Injection)
func NewNotificationRepository(db *gorm.DB) INotificationRepository {
	return &notificationRepository{DB: db}
}

func (r *notificationRepository) BeginTransaction() (*gorm.DB, error) {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (r *notificationRepository) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *notificationRepository) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}


func (r *notificationRepository) Save(tx *gorm.DB, input Notification) error {	
	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	

	if err := tx.WithContext(ctx).Create(&input).Error; err != nil {
		return err
	}
	return nil
}


func (r *notificationRepository) GetByUserID(tx *gorm.DB, userID int) ([]Notification, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var notifications []Notification
    return notifications, tx.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Find(&notifications).Error
}