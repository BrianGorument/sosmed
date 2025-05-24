package posts

import (
	"context"
	"fmt"
	"sosmed/src/interactions"
	"time"

	"gorm.io/gorm"
)

// userRepository struct
type postRepository struct {
	db *gorm.DB
}

// NewUserRepository (Dependency Injection)
func NewPostRepository(DB *gorm.DB) IPostRepository {
	return &postRepository{db: DB}
}

func (r *postRepository) BeginTransaction() (*gorm.DB, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (r *postRepository) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *postRepository) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *postRepository) InsertPosting(tx *gorm.DB, input Post_Content) (*PostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use GORM with context
	if err := tx.WithContext(ctx).Create(&input).Error; err != nil {
		return nil, err
	}

	return &PostResponse{Post_ID: input.ID , UserID: input.UserID}, nil
}


func (r *postRepository) InsertLikesTable(tx *gorm.DB, input interactions.Likes) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use GORM with context
	if err := tx.WithContext(ctx).Create(&input).Error; err != nil {
		return err
	}

	return nil	
}

func (r *postRepository) FindAll(filter GetAllPostsFilterRequest, user UserData) (*GetAllPostsResponse, error) {
	//var posts []*Post_Content
	var postResponse []*PagiPostsRespone
	var totalCount int64
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	
	db := r.db.WithContext(ctx).Table("post_content")
	db.Select("post_content.id , post_content.user_id,(u.username) as poster_name, (post_content.title) as post_title , post_content.media , post_content.like_count , post_content.category_id , post_content.created_at , post_content.updated_at").
	Joins("inner join users u on post_content.user_id  = u.id")
	
	
	if filter.PostID != 0 || len(filter.Title) != 0 {
		db = db.Where("post_content.id = ?", filter.PostID)
	}

	if filter.Title != "" {
		db = db.Where("post_content.title LIKE ?", fmt.Sprintf("%%%s%%", filter.Title))
	}

	if filter.ByUserName != "" {
		likePattern := "%" + filter.ByUserName + "%"
		db = db.Where("u.first_name like ? or u.username like ?", likePattern , likePattern)
	}
	
	if filter.ByUserID != 0 {
		db = db.Where("u.id = ?", filter.ByUserID)
	}
	
	// Count the total number of posts (for pagination purposes)
	err := db.Count(&totalCount).Error
	if err != nil {
		return nil, err
	}
	
	offset := (filter.Page - 1) * filter.Limit
	err = db.Limit(filter.Limit).Offset(offset).Find(&postResponse).Error
	if err != nil {
		return nil, err
	}	
	
	var postDTOs []*PagiPostsRespone
	for _, post := range postResponse {
		postDTO := &PagiPostsRespone{
			ID:          post.ID,
			UserID:      post.UserID,
			PosterName:  post.PosterName,  
			Post_Title:  post.Post_Title,
			Media:       post.Media,
			LikeCount:   post.LikeCount,
			CategoryID:  post.CategoryID,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt, 
		}
		postDTOs = append(postDTOs, postDTO)
	}
	
	return &GetAllPostsResponse{
		Posts:      postDTOs,
		Limit:      filter.Limit,
		Page:       filter.Page,
		TotalCount: int(totalCount),
	}, nil
}