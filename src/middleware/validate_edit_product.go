package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/scottdiddy/sal-app/src/constants"
	"github.com/scottdiddy/sal-app/src/models"
	"github.com/scottdiddy/sal-app/src/utils"
)

func ValidateEditProductDetails(c *fiber.Ctx) error {
	var updateFields models.UpdateProductDTO
	err := c.BodyParser(&updateFields)
	if err != nil {
		msg := utils.ResponseMessage(fmt.Sprintf("Error parsing request body, %s", err.Error()), nil)
		return c.Status(400).JSON(msg)
	}
	err = utils.ValidateStructBody(&updateFields)
	if err != nil {
		msg := utils.ResponseMessage(err.Error(), nil)
		return c.Status(400).JSON(msg)
	}
	c.Locals(constants.EDIT_PRODUCT_ROUTE, &updateFields)
	return c.Next()

}
