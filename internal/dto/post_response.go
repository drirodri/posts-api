package dto
import (
	"posts-api/internal/models"
	"time"
)
type UserData struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
type PostResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int64     `json:"author_id"`
	Author    *UserData `json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type PostListResponse struct {
	Posts      []PostResponse `json:"posts"`
	TotalCount int64          `json:"total_count"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalPages int            `json:"total_pages"`
}
func (pr *PostResponse) FromModel(post *models.Post) {
	pr.ID = post.ID
	pr.Title = post.Title
	pr.Content = post.Content
	pr.AuthorID = post.AuthorID
	pr.CreatedAt = post.CreatedAt
	pr.UpdatedAt = post.UpdatedAt
	
	if post.AuthorName != "" && post.AuthorEmail != "" {
		pr.Author = &UserData{
			ID:       post.AuthorID,
			Username: post.AuthorName,
			Email:    post.AuthorEmail,
		}
	}
}
func (pr *PostResponse) FromModelWithUser(post *models.Post, user *UserData) {
	pr.FromModel(post)
	pr.Author = user
}
func NewPostResponse(post *models.Post) *PostResponse {
	response := &PostResponse{}
	response.FromModel(post)
	return response
}
func NewPostResponseWithUser(post *models.Post, user *UserData) *PostResponse {
	response := &PostResponse{}
	response.FromModelWithUser(post, user)
	return response
}
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
