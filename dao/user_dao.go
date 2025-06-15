package dao

import "gorm.io/gorm"

// User model represents the users table
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
}
