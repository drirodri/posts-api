// post_repository.go - Post data access layer
package repository

import (
	"posts-api/internal/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *models.Post) error
	GetPostByID(id int64) (*models.Post, error)
	GetAllPosts() ([]*models.Post, error)
	UpdatePost(post *models.Post) error
	DeletePost(id int64) error
	GetByAuthorID(authorID int64) ([]*models.Post, error)
	GetTotalPosts() (int64, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) CreatePost(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) GetPostByID(id int64) (*models.Post, error) {
	var post models.Post
	err := r.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) GetAllPosts() ([]*models.Post, error) {
	var posts []*models.Post
	err := r.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) DeletePost(id int64) error {
	return r.db.Delete(&models.Post{}, id).Error
}

func (r *postRepository) UpdatePost(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) GetByAuthorID(authorID int64) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.db.Where("author_id = ?", authorID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) GetTotalPosts() (int64, error) {
	var count int64
	err := r.db.Model(&models.Post{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}




