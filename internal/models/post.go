// post.go - Post model definition
package models 

type Post struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string `json:"title" gorm:"not null"`
	Content   string `json:"content" gorm:"not null"`
	AuthorID  int64  `json:"author_id" gorm:"not null"`
	Author    User   `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}