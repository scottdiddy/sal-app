package repository

import "github.com/scottdiddy/sal-app/src/models"

var (
	ProductData  map[string]models.Product = map[string]models.Product{} //Maps the SKUID to the product details
	MerchantData map[string][]string       = map[string][]string{}       //Maps the MerchantID to the merchant's products' SKUID
)
