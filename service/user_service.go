package service

import (
	"time"

	"github.com/emanuelhardwell/go-user/dao"
	"github.com/emanuelhardwell/go-user/dto"
	"github.com/emanuelhardwell/go-user/utils"
)

// Create creates a new user
type User = dao.User

func Create(input dto.UserCreateDTO) (*dto.UserResponseDTO, error) {
	if err := utils.ValidateStruct(input); err != nil {
		return nil, err
	}
	hashed, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &User{Name: input.Name, Email: input.Email, Password: hashed}

	if err := dao.Create(user); err != nil {
		return nil, err
	}

	return &dto.UserResponseDTO{ID: user.ID, Name: user.Name, Email: user.Email, CreatedAt: time.Now()}, nil
}
