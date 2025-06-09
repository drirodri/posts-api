// user.go - User reference model for Users API integration
package models

import "time"
type User struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}