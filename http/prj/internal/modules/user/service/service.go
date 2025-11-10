package service

import (
	"context"
	"errors"
	"prj/internal/modules/user/model"
	"prj/internal/modules/user/repository"
	"prj/pkg/utils"
)

type Service struct {
	repo       repository.UserRepositoryInterface
	jwtService JWTService
}

func NewService() *Service {
	return &Service{
		repo:       repository.New(),
		jwtService: NewJWTService(),
	}
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

func (s *Service) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if !utils.CheckPassword(user.Password, password) {
		return "", errors.New("invalid username or password")
	}

	// Generate JWT token for this user
	token, err := s.jwtService.GenerateToken(user.Username)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
