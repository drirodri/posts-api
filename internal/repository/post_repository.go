// post_repository.go - Post data access layer
package repository

import (
	"posts-api/internal/models"
)
type PostRepository interface {
	CreatePost(post *models.Post) error
	GetPostByID(id int64) (*models.Post, error)
	GetAllPosts() ([]*models.Post, error)
	UpdatePost(post *models.Post) error
	DeletePost(id int64) error
	GetByAuthorID(author string) ([]*models.Post, error)
	GetTotalPosts() (int64, error)
}