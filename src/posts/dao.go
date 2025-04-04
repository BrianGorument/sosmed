package users

import "time"

type Post_Content struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    string      `gorm:"unique;not null" json:"user_id"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	Image 	  string    `gorm:"not null" json:"image"` 
	Media	  string    `gorm:"not null" json:"media"`
	LikeCount int       `gorm:"not null" json:"like_count"`
	CategoryId int       `gorm:"not null" json:"category_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func (Post_Content) TableName() string {
	return "post_content"
}
