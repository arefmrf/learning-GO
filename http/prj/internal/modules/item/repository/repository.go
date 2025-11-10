package repository

import (
	"context"
	"gorm.io/gorm"
	"prj/internal/modules/item/model"
	"prj/pkg/database"
)

type ItemRepository struct {
	DB *gorm.DB
}

func New() *ItemRepository {
	return &ItemRepository{DB: database.Connection()}
}

func (r *ItemRepository) List(ctx context.Context, limit, offset int) ([]model.Item, error) {
	var items []model.Item
	if err := r.DB.WithContext(ctx).Limit(limit).Offset(offset).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemRepository) Create(ctx context.Context, input *model.CreateItem) (*model.Item, error) {
	dbItem := model.Item{
		Name:  input.Name,
		Price: input.Price,
	}

	if err := r.DB.WithContext(ctx).Create(&dbItem).Error; err != nil {
		return nil, err
	}

	return &dbItem, nil
}
