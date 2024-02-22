package services

import (
	"errors"

	"github.com/scottdiddy/sal-app/src/models"
	"github.com/scottdiddy/sal-app/src/repository"
)

func DisplayProducts(merchantID string) ([]models.Product, error){
	//Checks if the merchantID is empty or if the merchant exists
	merchantProductSkuids, merchantExists := repository.MerchantData[merchantID]
	if (merchantID == ":merchantID") || (!merchantExists) {
		return []models.Product{}, errors.New("no merchant ID provided or Merchant does not exist")
	}
	var merchantProducts []models.Product

	//ranges over the skuids of a merchant. Uses those skuids to get all associated products and 
	//save them to slice.
	for _, skuid := range merchantProductSkuids {
		merchantProducts = append(merchantProducts, repository.ProductData[skuid])
	}
	return merchantProducts, nil

}
