package services

import (
	"errors"
	"time"

	"github.com/scottdiddy/sal-app/src/models"
	"github.com/scottdiddy/sal-app/src/repository"
)

func CreateNewProduct(newProduct *models.Product, merchantID string) error {
	newProduct.DateAdded = time.Now()
	newProduct.DateUpdated = time.Now()

	//if the product already exists, return 
	if _, exists := repository.ProductData[newProduct.SKUID]; exists  {
		return errors.New("product already exists")
	}
	//else save the product itself to the product data structure using it's skuid as key
	repository.ProductData[newProduct.SKUID] = *newProduct

	//saves the product's skuid to the string of skuids of the merchant
	repository.MerchantData[merchantID]= append(repository.MerchantData[merchantID], newProduct.SKUID)
	return nil
}