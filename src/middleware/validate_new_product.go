package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/scottdiddy/sal-app/src/constants"
	"github.com/scottdiddy/sal-app/src/models"
	"github.com/scottdiddy/sal-app/src/utils"
)

func ValidateNewProduct(c *fiber.Ctx) error {
	var reqBody models.ProductDTO
	err := c.BodyParser(&reqBody)
	if err != nil {
		msg := utils.ResponseMessage(false, fmt.Sprintf("Error parsing request body, %s", err.Error()), nil)
		return c.Status(400).JSON(msg)
	}
	err = utils.ValidateStructBody(&reqBody)
	if err != nil {
		msg := utils.ResponseMessage(false, err.Error(), nil)
		return c.Status(400).JSON(msg)
	}
	newProduct := models.Product{
		ProductDTO: &reqBody,
	}
	c.Locals(constants.CREATE_PRODUCT_ROUTE, &newProduct)
	return c.Next()
}
