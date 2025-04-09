package interactions

import "time"

type Comments struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    uint      `gorm:"unique;not null" json:"post_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Comment   string    `json:"comment"` 
	Media	  string    `json:"media"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (Comments) TableName() string {
	return "comments"
}

type Likes struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    uint      `gorm:"unique;not null" json:"post_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Type	  string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (Likes) TableName() string {
	return "likes"
}
