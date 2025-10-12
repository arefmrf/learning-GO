package repository

import (
	"context"
	"gorm.io/gorm"
	"prj/internal/modules/user/model"
	"prj/pkg/database"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{DB: database.Connection()}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return &user, err
}
