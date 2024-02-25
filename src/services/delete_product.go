package services

import (
	"errors"

	"github.com/scottdiddy/sal-app/src/repository"
	"github.com/scottdiddy/sal-app/src/utils"
)

func DeleteProduct(merchantID string, skuidToDelete string) error {
	merchantProductSkuids, err := utils.CheckMerchant(merchantID)
	if err != nil {
		return err
	}

	productIndex := 0
	productMatch := false

	//Checks if the supplied skuid matches the any skuid of the merchant's products
	for i, skuidofMerchant := range merchantProductSkuids {
		if skuidToDelete == skuidofMerchant  {
			productIndex = i
			productMatch = true
			break
		}
	}

	//if product matches, delete the product from our products data structure
	//also delete the product's skuid from the slice of the skuids of a merchant's products
	//and save the slice
	if productMatch {
		delete(repository.ProductData, skuidToDelete)
		merchantProductSkuids = append(merchantProductSkuids[:productIndex], merchantProductSkuids[productIndex+1:]...)
		repository.MerchantData[merchantID] = merchantProductSkuids
		return nil
	}
	//if product does not exist, throw error
	return errors.New("supplied skuid does not match skuid of merchant's product")

}