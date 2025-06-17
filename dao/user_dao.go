package dao

import (
	"github.com/emanuelhardwell/go-user/config"
	"github.com/emanuelhardwell/go-user/model"
)

// Create inserts a new user record
type User = model.User

func Create(user *User) error {
	return config.DB.Create(user).Error
}
