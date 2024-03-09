package main

import (
	"log"

	"github.com/Vishal21121/go-auth-mysql.git/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
