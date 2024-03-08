package main

import (
	"log"

	"github.com/Vishal21121/go-auth-mysql.git/controllers"
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

	userRoutes := app.Group("/api/v1/users")
	userRoutes.Post("/register", controllers.RegisterUser)
	userRoutes.Post("/login", controllers.LoginUser)

	log.Printf("Server is running on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
