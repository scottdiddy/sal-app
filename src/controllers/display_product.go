package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/scottdiddy/sal-app/src/services"
	"github.com/scottdiddy/sal-app/src/utils"
)

func DisplayProducts(c *fiber.Ctx) error {
	//gets the merchantID from the header
	merchantID := fmt.Sprint(c.Params("merchantID"))

	merchantProducts, err := services.DisplayProducts(merchantID)
	if err != nil {
		msg := utils.ResponseMessage(err.Error(), nil)
		return c.Status(404).JSON(msg)
	}

	// fmt.Printf("\nIn Display product Controller lower\n%v\n%v\n%v\n", merchantID, repository.MerchantData, repository.ProductData)
	msg := utils.ResponseMessage("Display products successful", merchantProducts)
	return c.Status(200).JSON(msg)
}
