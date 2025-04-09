package interactions

import "time"

type Post_Interactions struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    uint      `gorm:"unique;not null" json:"post_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Type  	  string    `gorm:"not null" json:"type"`
	Comment   string    `json:"comment"` 
	Media	  string    `json:"media"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (Post_Interactions) TableName() string {
	return "post_interactions"
}
