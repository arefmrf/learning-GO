package dto

type CreateItemDTO struct {
	Title string  `json:"title" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type UpdateItemDTO struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}
