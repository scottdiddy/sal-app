package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/scottdiddy/sal-app/src/constants"
)

func RegisterRoutes(appGroup fiber.Router) {
	appGroup.Post(constants.CREATE_PRODUCT_ROUTE, CreateProduct)
	appGroup.Get(constants.DISPLAY_PRODUCT_ROUTE, DisplayProducts)
	appGroup.Patch(constants.EDIT_PRODUCT_ROUTE, EditProduct)
	appGroup.Delete(constants.DELETE_PRODUCT_ROUTE, DeleteProduct)
}