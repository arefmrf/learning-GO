package service

import (
	"main/internal/modules/item/model"
	"main/internal/modules/item/repository"
)

type ItemService interface {
	Create(item *model.Item) error
	GetAll() ([]model.Item, error)
	GetByID(id uint) (*model.Item, error)
	Update(item *model.Item) error
	Delete(id uint) error
}

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{repo}
}

func (s *itemService) Create(item *model.Item) error {
	return s.repo.Create(item)
}

func (s *itemService) GetAll() ([]model.Item, error) {
	return s.repo.FindAll()
}

func (s *itemService) GetByID(id uint) (*model.Item, error) {
	return s.repo.FindByID(id)
}

func (s *itemService) Update(item *model.Item) error {
	return s.repo.Update(item)
}

func (s *itemService) Delete(id uint) error {
	return s.repo.Delete(id)
}
