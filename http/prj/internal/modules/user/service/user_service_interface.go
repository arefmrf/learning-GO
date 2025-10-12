package service

import (
	"context"
	"prj/internal/modules/user/model"
)

type UserServiceInterface interface {
	Register(ctx context.Context, username, password string) error
	Login(ctx context.Context, username, password string) (*model.User, error)
}
