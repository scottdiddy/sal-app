package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/scottdiddy/sal-app/src/repository"
	"github.com/scottdiddy/sal-app/src/services"
	"github.com/scottdiddy/sal-app/src/utils"
)

func DeleteProduct(c *fiber.Ctx) error {
	//gets the merchantID from the header
	merchantID := fmt.Sprint(c.Params("merchantID"))
	
	//Checks if the merchantID is empty or if the merchant exists
	_, merchantExists := repository.MerchantData[merchantID]
	if (merchantID == ":merchantID") || ( !merchantExists ) {
		msg := utils.ResponseMessage("No merchant ID provided or Merchant does not exist", nil)
		return c.Status(400).JSON(msg)
	}

	//gets the skuid of the product to be delete from the header
	skuidToDelete := fmt.Sprint(c.Query("sku_id"))

	//checks if the skuid is empty
	if (skuidToDelete == ""){
		msg := utils.ResponseMessage("No product sku_id provided", nil)
		return c.Status(400).JSON(msg)
	}
	
	err := services.DeleteProduct(merchantID, skuidToDelete)
	if err != nil {
		msg := utils.ResponseMessage(err.Error(), nil)
		return c.Status(400).JSON(msg)
	}
	
	msg := utils.ResponseMessage("Successfully deleted product", fiber.Map{
		"SKUID": skuidToDelete,
	})
	return c.Status(400).JSON(msg)


}