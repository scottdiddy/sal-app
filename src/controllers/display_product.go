package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/scottdiddy/sal-app/src/repository"
	"github.com/scottdiddy/sal-app/src/services"
	"github.com/scottdiddy/sal-app/src/utils"
)

func DisplayProducts(c *fiber.Ctx) error {
	//gets the merchantID from the header
	merchantID := fmt.Sprint(c.Params("merchantID"))

	//Checks if the merchantID is empty or if the merchant exists
	merchantProductSkuids, merchantExists := repository.MerchantData[merchantID]
	if (merchantID == ":merchantID") || (!merchantExists) {
		msg := utils.ResponseMessage("No merchant ID provided or Merchant does not exist", nil)
		return c.Status(404).JSON(msg)
	}
	merchantProducts := services.DisplayProducts(merchantProductSkuids)

	// fmt.Printf("\nIn Display product Controller lower\n%v\n%v\n%v\n", merchantID, repository.MerchantData, repository.ProductData)
	msg := utils.ResponseMessage("Display products successful", merchantProducts)
	return c.Status(200).JSON(msg)
}
