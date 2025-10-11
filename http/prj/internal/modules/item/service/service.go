package service

import (
	"context"
	"prj/internal/modules/item/model"
	"prj/internal/modules/item/repository"
)

type Service struct {
	repo repository.ItemRepositoryInterface
}

func NewService() *Service {
	return &Service{repo: repository.New()}
}

func (s *Service) List(ctx context.Context, limit, offset int) ([]model.Item, error) {
	return s.repo.List(ctx, limit, offset)
}
