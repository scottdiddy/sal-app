package services

import (
	"errors"
	"time"

	"github.com/scottdiddy/sal-app/src/models"
	"github.com/scottdiddy/sal-app/src/repository"
	"github.com/scottdiddy/sal-app/src/utils"
)

func EditProduct(updateFields *models.UpdateProductDTO, merchantProductSkuids []string) (models.Product, error) {
	productMatch := false

	//Checks if the supplied skuid matches the any skuid of the merchant's products
	for _, skuid := range merchantProductSkuids {
		if skuid == updateFields.SKUID {
			productMatch = true
			break
		}
	}

	//if product matches, get the product and update the product's fields
	if productMatch {
		productToUpdate := repository.ProductData[updateFields.SKUID]
		utils.UpdateStructFields(productToUpdate.ProductDTO, updateFields)
		productToUpdate.DateUpdated = time.Now()
		repository.ProductData[updateFields.SKUID] = productToUpdate //save the updated product
		return productToUpdate, nil
	}

	//if product does not match, throw error
	return models.Product{}, errors.New("supplied skuid does not match skuid of merchant's product")
}
