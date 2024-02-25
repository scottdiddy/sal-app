package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/scottdiddy/sal-app/src/constants"
	"github.com/scottdiddy/sal-app/src/models"
	"github.com/scottdiddy/sal-app/src/services"
	"github.com/scottdiddy/sal-app/src/utils"
)

func CreateProduct(c *fiber.Ctx) error {
	//gets the newproduct saved in the context by validate middleware
	newProduct, _ := c.Locals(constants.CREATE_PRODUCT_ROUTE).(*models.Product)

	//gets the merchantID from the header. Uses fmt.Sprint to avoid any leading slash (/)
	merchantID := fmt.Sprint(c.Params("merchantID"))

	//checks if the merchantId is empty
	if merchantID == ":merchantID" {
		msg := utils.ResponseMessage(false, "No merchant ID provided", nil)
		return c.Status(400).JSON(msg)
	}
	err := services.CreateNewProduct(newProduct, merchantID)
	if err != nil {
		msg := utils.ResponseMessage(false, err.Error(), nil)
		return c.Status(400).JSON(msg)
	}

	msg := utils.ResponseMessage(true, "Successfully created product", nil)
	return c.Status(200).JSON(msg)

}
