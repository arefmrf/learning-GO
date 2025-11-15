package repository

import (
	"gorm.io/gorm"
	"main/internal/modules/item/model"
)

type ItemRepository interface {
	Create(item *model.Item) error
	FindAll() ([]model.Item, error)
	FindByID(id uint) (*model.Item, error)
	Update(item *model.Item) error
	Delete(id uint) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) Create(item *model.Item) error {
	return r.db.Create(item).Error
}

func (r *itemRepository) FindAll() ([]model.Item, error) {
	var items []model.Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r *itemRepository) FindByID(id uint) (*model.Item, error) {
	var item model.Item
	err := r.db.First(&item, id).Error
	return &item, err
}

func (r *itemRepository) Update(item *model.Item) error {
	return r.db.Save(item).Error
}

func (r *itemRepository) Delete(id uint) error {
	return r.db.Delete(&model.Item{}, id).Error
}
