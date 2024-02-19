package repository

import "github.com/scottdiddy/sal-app/src/models"

var (
	ProductData  = make(map[string]models.Product)   //Maps the SKUID to the product details
	MerchantData = make(map[string][]string) //Maps the MerchantID to the merchant's products' SKUID
)
