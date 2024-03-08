package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	//TODO: Implement user registration
	fmt.Println(string(c.Body()))
	return c.SendStatus(200)
}

func LoginUser(c *fiber.Ctx) error {
	//TODO: Implement user login
	fmt.Println(string(c.Body()))
	return c.SendStatus(200)
}
