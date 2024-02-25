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
		msg := utils.ResponseMessage(false, err.Error(), nil)
		return c.Status(400).JSON(msg)
	}

	msg := utils.ResponseMessage(true, "Display products successful", merchantProducts)
	return c.Status(200).JSON(msg)
}
