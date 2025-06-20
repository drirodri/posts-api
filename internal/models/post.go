package models
import "time"
type Post struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"not null"`
	Content     string    `json:"content" gorm:"not null"`
	AuthorID    int64     `json:"author_id" gorm:"not null"`
	AuthorName  string    `json:"author_name" gorm:"not null"`
	AuthorEmail string    `json:"author_email" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
