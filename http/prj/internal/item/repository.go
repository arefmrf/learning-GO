package item

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/gorm"
	"prj/pkg/database"
)

type ItemRepository struct {
	DB *gorm.DB
}

func NewRepository() *ItemRepository {
	return &ItemRepository{
		DB: database.Connection(),
	}
}

func (r *ItemRepository) List(ctx context.Context, limit, offset int) ([]Item, error) {
	query := fmt.Sprintf(`SELECT id, name, price FROM items ORDER BY id LIMIT %d OFFSET %d`, limit, offset)

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var it Item
		if err := rows.Scan(&it.ID, &it.Name, &it.Price); err != nil {
			return nil, err
		}
		items = append(items, it)
	}

	return items, nil
}
