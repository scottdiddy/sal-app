package models

import "time"

type ProductDTO struct {
	SKUID       string  `json:"sku_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}
type Product struct {
	*ProductDTO
	DateAdded   time.Time
	DateUpdated time.Time
}
