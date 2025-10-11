package model

type Item struct {
	ID    int64  `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}
