package posts

import (
	"sosmed/src/interactions"

	"gorm.io/gorm"
)

type IPostService interface {
	CreatePosting(req CreatePostRequest, users UserData) (*PostResponse, error)
	GetAllPosts(filter GetAllPostsFilterRequest , user UserData) (*GetAllPostsResponse , error)
}

type IPostRepository interface {
	BeginTransaction() (*gorm.DB, error)               // Untuk memulai transaksi
	CommitTransaction(tx *gorm.DB) error               // Untuk commit transaksi
	RollbackTransaction(tx *gorm.DB) error 
	InsertPosting(tx *gorm.DB, input Post_Content) (*PostResponse, error)
	InsertLikesTable(tx *gorm.DB, input interactions.Likes) error
	FindAll(filter GetAllPostsFilterRequest, user UserData) (*GetAllPostsResponse, error)
}
