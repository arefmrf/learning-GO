package model

import "time"

type Item struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
