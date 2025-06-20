package dto
import (
	"posts-api/internal/models"
	"time"
)
type CreatePostRequest struct {
    Title   string `json:"title" validate:"required,min=1,max=255"`
    Content string `json:"content" validate:"required,min=1,max=10000"`
}
func (req *CreatePostRequest) ToModel(authorID int64) *models.Post {
    return &models.Post{
        Title:     req.Title,
        Content:   req.Content,
        AuthorID:  authorID,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
}
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
func (req *CreatePostRequest) Validate() error {
    return nil
}
