package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.SendStatus(200)
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Server health is ok",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
