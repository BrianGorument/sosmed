package notifications

import "time"

type Notification struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    Message   string    `json:"message"`
    Type      string    `json:"type"`
    CreatedAt time.Time `json:"created_at"`
    IsRead    bool      `json:"is_read"`
}

func (Notification) TableName() string {
	return "notifications"
}