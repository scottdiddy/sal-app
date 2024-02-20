package models

type UpdateProductDTO struct {
	SKUID       string  `json:"sku_id" validate:"required,validSKUID"` //validSKUID checks if the product exists
	Name        string  `json:"name" validate:"omitempty"`
	Description string  `json:"description" validate:"omitempty"`
	Price       float64 `json:"price" validate:"omitempty"`
}
