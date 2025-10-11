package repository

import (
	"context"
	"prj/internal/modules/item/model"
)

type ItemRepositoryInterface interface {
	List(ctx context.Context, limit, offset int) ([]model.Item, error)
}
