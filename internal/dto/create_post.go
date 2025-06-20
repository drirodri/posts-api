// create_post.go - DTO for creating posts
package dto

import (
	"posts-api/internal/models"
	"time"
)

// CreatePostRequest represents the request body for creating a new post
type CreatePostRequest struct {
    Title   string `json:"title" validate:"required,min=1,max=255"`
    Content string `json:"content" validate:"required,min=1,max=10000"`
}

// ToModel converts the DTO to a Post model
func (req *CreatePostRequest) ToModel(authorID int64) *models.Post {
    return &models.Post{
        Title:     req.Title,
        Content:   req.Content,
        AuthorID:  authorID,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
}

// ToModelWithAuthor converts the DTO to a Post model with author information
func (req *CreatePostRequest) ToModelWithAuthor(authorID int64, authorName, authorEmail string) *models.Post {
    return &models.Post{
        Title:       req.Title,
        Content:     req.Content,
        AuthorID:    authorID,
        AuthorName:  authorName,
        AuthorEmail: authorEmail,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
}

// Validate performs custom business validation
func (req *CreatePostRequest) Validate() error {
    // Add any additional business validation here
    // For example: check for prohibited words, etc.
    return nil
}