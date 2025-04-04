package users

type CreatePostRequest struct {
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	Image   string `gorm:"not null" json:"image"`
	Media   string `gorm:"not null" json:"media"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required"`
}

type PostResponse struct {
	Post_ID uint   `json:"Post_id"`
	UserID  string `json:"user_id"`
}

type UserData struct {
	UserId    float64 `json:"user_id"`
	Username  string  `json:"user_name"`
	UserEmail string  `json:"user_email"`
}