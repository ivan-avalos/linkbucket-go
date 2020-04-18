package database

import "github.com/jinzhu/gorm"

// PasswordReset represents a password reset
type PasswordReset struct {
	gorm.Model
	email string
	token string
}
