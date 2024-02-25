package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/scottdiddy/sal-app/src/constants"
)

func RegisterRoutes(appGroup fiber.Router) {
	appGroup.Post(constants.CREATE_PRODUCT_ROUTE, ValidateNewProduct)
	appGroup.Patch(constants.EDIT_PRODUCT_ROUTE, ValidateEditProductDetails)
}