package models

import "time"

type Product struct {
	SKUID       string  `json:"sku_id" validate:"required,validSKUID"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	DateAdded   time.Time
}
