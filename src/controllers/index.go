package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/scottdiddy/sal-app/src/constants"
)

func RegisterRoutes(apiGroup fiber.Router) {
	apiGroup.Post(constants.CREATE_PRODUCT_ROUTE, CreateProduct)
	apiGroup.Get(constants.DISPLAY_PRODUCT_ROUTE, DisplayProducts)
	apiGroup.Patch(constants.EDIT_PRODUCT_ROUTE, EditProduct)
	apiGroup.Delete(constants.DELETE_PRODUCT_ROUTE, DeleteProduct)
}