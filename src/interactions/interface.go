package interactions

import "gorm.io/gorm"

type IInteractionService interface {
	InsertOrUpdateInteraction(req InteractRequest, user UserData) (*InteractResponse, error)
	DeleteCommentOrMedia(req DeleteCommentRequest, user UserData) (*InteractResponse, error)
}

type IInteractionRepository interface {
	BeginTransaction() (*gorm.DB, error)               // Untuk memulai transaksi
	CommitTransaction(tx *gorm.DB) error               // Untuk commit transaksi
	RollbackTransaction(tx *gorm.DB) error 
	InsertComment(tx *gorm.DB, input Comments) (*InteractResponse, error)
	UpdateLikesInteraction(tx *gorm.DB, input Likes) (*InteractResponse, error)
	DeleteCommentByID(tx *gorm.DB,comment Comments) (*InteractResponse, error)
	//DeleteMediaByID(tx *gorm.DB,comment Comments) (*InteractResponse, error)
	
}
