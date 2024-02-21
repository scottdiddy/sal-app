package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/scottdiddy/sal-app/src/constants"
	"github.com/scottdiddy/sal-app/src/models"
	"github.com/scottdiddy/sal-app/src/repository"
	"github.com/scottdiddy/sal-app/src/services"
	"github.com/scottdiddy/sal-app/src/utils"
)

func EditProduct(c *fiber.Ctx) error {
	//gets the updateFields saved in the context
	updateFields, _ := c.Locals(constants.EDIT_PRODUCT_ROUTE).(*models.UpdateProductDTO)
	
	//gets the merchantID from the header
	merchantID := fmt.Sprint(c.Params("merchantID"))

	//Checks if the merchantID is empty or if the merchant exists
	merchantProductSkuids, merchantExists := repository.MerchantData[merchantID]
	if (merchantID == ":merchantID") || ( !merchantExists ) {
		msg := utils.ResponseMessage("No merchant ID provided or Merchant does not exist", nil)
		return c.Status(400).JSON(msg)
	}
	updatedProduct, err := services.EditProduct(updateFields, merchantProductSkuids)
	if err != nil {
		msg := utils.ResponseMessage(err.Error(), nil)
		return c.Status(400).JSON(msg)
	}
	msg := utils.ResponseMessage("Successfully updated product", updatedProduct)
	return c.Status(200).JSON(msg)

}