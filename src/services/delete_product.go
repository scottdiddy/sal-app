package services

import (
	"errors"
	"github.com/scottdiddy/sal-app/src/repository"
)

func DeleteProduct(merchantID string, skuidToDelete string) error {
	//get all skuids of the products of the merchant
	merchantProductSkuids := repository.MerchantData[merchantID]

	productIndex := 0
	productExists := false

	//Checks if the supplied skuid matches the any skuid of the merchant's products
	for i, skuidofMerchant := range merchantProductSkuids {
		if skuidofMerchant == skuidToDelete {
			productIndex = i
			productExists = true
			break
		}
	}

	//if product exists, delete the product from our products data structure
	//also delete the product skuid from the slice of the skuids of a merchant's products
	//and save the slice
	if productExists {
		delete(repository.ProductData, skuidToDelete)
		merchantProductSkuids = append(merchantProductSkuids[:productIndex], merchantProductSkuids[productIndex+1:]...)
		repository.MerchantData[merchantID] = merchantProductSkuids
	}
	//if product does not exist, throw error
	return errors.New("supplied skuid does not match skuid of merchant's product")

}