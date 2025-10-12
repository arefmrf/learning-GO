package service

import (
	"context"
	"errors"
	"prj/internal/modules/user/model"
	"prj/internal/modules/user/repository"
	"prj/pkg/utils"
)

type Service struct {
	repo repository.UserRepositoryInterface
}

func NewService() *Service {
	return &Service{repo: repository.New()}
}

func (s *Service) Register(ctx context.Context, username, password string) error {
	user := &model.User{Username: username}
	hashedPass, err := utils.HashPassword(password)
	if err != nil || hashedPass == "" {
		return err
	}
	user.Password = hashedPass
	return s.repo.Create(ctx, user)
}

func (s *Service) Login(ctx context.Context, username, password string) (*model.User, error) {
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if !utils.CheckPassword(user.Password, password) {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}
