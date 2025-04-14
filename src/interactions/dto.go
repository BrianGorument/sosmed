package interactions

type InteractRequest struct {
	PostID  uint   `gorm:"unique;not null" json:"post_id"`
	Type    string `gorm:"not null" json:"type"`
	Comment string `json:"comment"`
	Media   string `json:"media"`
}

type InteractResponse struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID  uint   `gorm:"unique;not null" json:"post_id"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	Comment string `json:"comment"`
	Media   string `json:"media"`
}

type DeleteCommentRequest struct {
	ID     uint `gorm:"unique;not null" json:"id"`
	PostID uint `gorm:"unique;not null" json:"post_id"`
}
type UserData struct {
	UserId    uint   `json:"user_id"`
	Username  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}