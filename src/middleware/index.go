package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/scottdiddy/sal-app/src/constants"
)

func RegisterRoutes(apiGroup fiber.Router) {
	apiGroup.Post(constants.CREATE_PRODUCT_ROUTE, ValidateNewProduct)
	apiGroup.Patch(constants.EDIT_PRODUCT_ROUTE, ValidateEditProductDetails)
}