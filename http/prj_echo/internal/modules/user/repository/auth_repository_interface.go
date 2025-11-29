package repository

import (
	"main/internal/modules/user/dto"
)

type UserRepositoryInterface interface {
	Create(user *dto.RegisterDTO) error
}
