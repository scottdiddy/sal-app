package services

import (
	"github.com/scottdiddy/sal-app/src/models"
	"github.com/scottdiddy/sal-app/src/repository"
	"github.com/scottdiddy/sal-app/src/utils"
)

func DisplayProducts(merchantID string) ([]models.Product, error){
	merchantProductSkuids, err := utils.CheckMerchant(merchantID)
	if err != nil {
		return nil, err
	}
	var merchantProducts []models.Product
	//ranges over the skuids of a merchant. Uses those skuids to get all associated products and 
	//save them to the slice.
	for _, skuid := range merchantProductSkuids {
		merchantProducts = append(merchantProducts, repository.ProductData[skuid])
	}
	return merchantProducts, nil

}
