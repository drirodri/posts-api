// user.go - User reference model for Users API integration
package models
type User struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}