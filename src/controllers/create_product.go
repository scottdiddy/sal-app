package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/scottdiddy/sal-app/src/constants"
	"github.com/scottdiddy/sal-app/src/models"
	"github.com/scottdiddy/sal-app/src/services"
	"github.com/scottdiddy/sal-app/src/utils"
)

func CreateProduct(c *fiber.Ctx) error {
	//gets the newproduct saved in the context
	newProduct, _ := c.Locals(constants.CREATE_PRODUCT_ROUTE).(*models.Product)
	
	//gets the merchantID from the header
	merchantID := c.Params("merchantID") 

	//checks if the merchantId is empty
	if merchantID == "" {
		msg := utils.ResponseMessage("No merchant ID provided", nil)
		return c.Status(404).JSON(msg)
	}
	services.CreateNewProduct(newProduct, merchantID)
	msg := utils.ResponseMessage("Successfully created product", nil)
	return c.Status(200).JSON(msg)

}
