package services

import (
	"github.com/scottdiddy/sal-app/src/models"
	"github.com/scottdiddy/sal-app/src/repository"
)

func DisplayProducts(merchantProductSkuids []string) []models.Product {
	var merchantProducts []models.Product

	//ranges over the skuids of a merchant. Uses those skuids to get all associated products and 
	//save them to slice.
	for _, skuid := range merchantProductSkuids {
		merchantProducts = append(merchantProducts, repository.ProductData[skuid])
	}
	return merchantProducts

}
