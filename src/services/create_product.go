package services

import (
	"errors"
	"time"

	"github.com/scottdiddy/sal-app/src/models"
	"github.com/scottdiddy/sal-app/src/repository"
)

func CreateNewProduct(newProduct *models.Product, merchantID string) error {
	//if the product already exists, return error
	if _, exists := repository.ProductData[newProduct.SKUID]; exists  {
		return errors.New("product already exists")
	}

	//else add the time the product was created and save the product
	//itself to the product data structure using it's skuid as key
	newProduct.DateAdded = time.Now()
	newProduct.DateUpdated = time.Now()
	repository.ProductData[newProduct.SKUID] = *newProduct

	//saves the product's skuid to the string of skuids of the merchant
	repository.MerchantData[merchantID]= append(repository.MerchantData[merchantID], newProduct.SKUID)

	return nil
}