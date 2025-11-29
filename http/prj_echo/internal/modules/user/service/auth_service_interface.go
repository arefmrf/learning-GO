package service

import (
	"main/internal/modules/user/dto"
)

type UserServiceInterface interface {
	Register(user *dto.RegisterDTO) error
}
