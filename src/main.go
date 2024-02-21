package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/scottdiddy/sal-app/src/controllers"
	"github.com/scottdiddy/sal-app/src/middleware"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	apiGroup := app.Group("/api")
	
	middleware.RegisterRoutes(apiGroup)
	controllers.RegisterRoutes(apiGroup)

	app.Use(func (c *fiber.Ctx) error  {
		return c.Status(404).JSON(fiber.Map{
			"Message": "Not found",
		})		
	})
	app.Listen(":3001")

}
