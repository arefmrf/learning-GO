package repository

import (
	"main/internal/modules/user/dto"
	"main/pkg/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{DB: database.Connection()}
}

func (r *UserRepository) Create(user *dto.RegisterDTO) error {
	r.DB.Create(user)
	return nil
}
