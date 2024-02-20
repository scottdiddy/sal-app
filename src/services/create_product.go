package services

import (
	"time"

	"github.com/scottdiddy/sal-app/src/models"
	"github.com/scottdiddy/sal-app/src/repository"
)

func CreateNewProduct(newProduct *models.Product, merchantID string){
	newProduct.DateAdded = time.Now()
	newProduct.DateUpdated = time.Now()

	//saves the product itself to the product data structure using it's skuid as key
	repository.ProductData[newProduct.SKUID] = *newProduct

	//saves the product's skuid to the string of skuids of the merchant
	repository.MerchantData[merchantID]= append(repository.MerchantData[merchantID], newProduct.SKUID)

}