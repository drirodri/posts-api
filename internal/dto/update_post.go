package dto

import (
	"posts-api/internal/models"
	"time"
)
type UpdatePostRequest struct {
	Title   *string `json:"title,omitempty" validate:"omitempty,min=1,max=255"`
	Content *string `json:"content,omitempty" validate:"omitempty,min=1,max=10000"`
}
func (req *UpdatePostRequest) UpdateModel(post *models.Post) {
	if req.Title != nil {
		post.Title = *req.Title
	}
	if req.Content != nil {
		post.Content = *req.Content
	}
	post.UpdatedAt = time.Now()
}
func (req *UpdatePostRequest) HasChanges() bool {
	return req.Title != nil || req.Content != nil
}
func (req *UpdatePostRequest) Validate() error {
	return nil
}
