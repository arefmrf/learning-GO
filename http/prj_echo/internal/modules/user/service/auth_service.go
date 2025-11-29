package service

import (
	"main/internal/modules/user/dto"
	"main/internal/modules/user/repository"
)

type Service struct {
	repo repository.UserRepositoryInterface
}

func NewService() *Service {
	return &Service{
		repo: repository.New(),
	}
}

func (s *Service) Register(user *dto.RegisterDTO) error {
	err := s.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
