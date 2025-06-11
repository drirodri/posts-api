// update_post.go - DTO for updating posts
package dto

import (
	"posts-api/internal/models"
	"time"
)

// UpdatePostRequest represents the request body for updating an existing post
type UpdatePostRequest struct {
	Title   *string `json:"title,omitempty" validate:"omitempty,min=1,max=255"`
	Content *string `json:"content,omitempty" validate:"omitempty,min=1,max=10000"`
}

// UpdateModel applies changes to an existing Post model
func (req *UpdatePostRequest) UpdateModel(post *models.Post) {
	if req.Title != nil {
		post.Title = *req.Title
	}
	if req.Content != nil {
		post.Content = *req.Content
	}
	post.UpdatedAt = time.Now()
}

// HasChanges returns true if any fields are set for update
func (req *UpdatePostRequest) HasChanges() bool {
	return req.Title != nil || req.Content != nil
}

// Validate performs custom business validation
func (req *UpdatePostRequest) Validate() error {
	// Add any additional business validation here
	return nil
}