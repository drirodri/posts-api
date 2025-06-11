// post_response.go - DTO for post responses
package dto

import (
	"posts-api/internal/models"
	"time"
)

// UserData represents user information from the Users API
type UserData struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// PostResponse represents a single post in API responses
type PostResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int64     `json:"author_id"`
	Author    *UserData `json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PostListResponse represents a list of posts with pagination
type PostListResponse struct {
	Posts      []PostResponse `json:"posts"`
	TotalCount int64          `json:"total_count"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalPages int            `json:"total_pages"`
}

// FromModel converts a Post model to PostResponse
func (pr *PostResponse) FromModel(post *models.Post) {
	pr.ID = post.ID
	pr.Title = post.Title
	pr.Content = post.Content
	pr.AuthorID = post.AuthorID
	pr.CreatedAt = post.CreatedAt
	pr.UpdatedAt = post.UpdatedAt
}

// FromModelWithUser converts a Post model to PostResponse and includes user data
func (pr *PostResponse) FromModelWithUser(post *models.Post, user *UserData) {
	pr.FromModel(post)
	pr.Author = user
}

// NewPostResponse creates a PostResponse from a Post model
func NewPostResponse(post *models.Post) *PostResponse {
	response := &PostResponse{}
	response.FromModel(post)
	return response
}

// NewPostResponseWithUser creates a PostResponse from a Post model with user data
func NewPostResponseWithUser(post *models.Post, user *UserData) *PostResponse {
	response := &PostResponse{}
	response.FromModelWithUser(post, user)
	return response
}

// NewPostListResponse creates a paginated list response
func NewPostListResponse(posts []*models.Post, totalCount int64, page, pageSize int) *PostListResponse {
	postResponses := make([]PostResponse, len(posts))
	for i, post := range posts {
		postResponses[i] = *NewPostResponse(post)
	}

	totalPages := int((totalCount + int64(pageSize) - 1) / int64(pageSize))

	return &PostListResponse{
		Posts:      postResponses,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}