package repository

import (
	"context"
	"prj/internal/modules/user/model"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *model.User) error
	FindByUsername(ctx context.Context, username string) (*model.User, error)
}
