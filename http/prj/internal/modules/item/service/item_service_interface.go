package service

import (
	"context"
	"prj/internal/modules/item/model"
)

type ItemServiceInterface interface {
	List(ctx context.Context, limit, offset int) ([]model.Item, error)
}
