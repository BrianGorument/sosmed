package posts

type CreatePostRequest struct {
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
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

type GetAllPostsFilterRequest struct {
	PostID     uint   `json:"post_id"`
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Title      string `json:"title"`
	ByUserName string `json:"username"`
}

type UserData struct {
	UserId    uint   `json:"user_id"`
	Username  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}

type PagiPostsRespone struct {
	ID         uint   `json:"id"`
	UserID     string `json:"user_id"`
	PosterName string `json:"poster_name"`
	Title      string `json:"title"`
	Media      string `json:"media"`
	LikeCount  int    `json:"like_count"`
	CategoryID int    `json:"category_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type GetAllPostsResponse struct {
	Posts      []*PagiPostsRespone `json:"posts"`
	Limit      int                 `json:"limit"`
	Page       int                 `json:"page"`
	TotalCount int                 `json:"total_count"`
}