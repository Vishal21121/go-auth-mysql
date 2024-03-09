package controllers

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type user struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Age       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

var db *gorm.DB

func DBSetter(database *gorm.DB) {
	db = database
	db.AutoMigrate(&user{})
}

func RegisterUser(c *fiber.Ctx) error {

	// userData receiving template
	type userData struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// converting the recived data in the following struct
	userDataReceived := new(userData)
	if err := c.BodyParser(userDataReceived); err != nil {
		log.Println(err)
	}

	// checking whether all the fields are provided
	if len(userDataReceived.Email) == 0 || len(userDataReceived.Name) == 0 || len(userDataReceived.Password) == 0 {
		c.SendStatus(400)
		return c.JSON(fiber.Map{
			"statusCode": 400,
			"data": fiber.Map{
				"success": false,
				"message": "Please provide all the details",
			},
		})
	}

	// if userDataReceived.

	// hashing the password
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(userDataReceived.Password), bcrypt.DefaultCost)
	if error != nil {
		c.SendStatus(500)
		return c.JSON(fiber.Map{
			"statusCode": 500,
			"data": fiber.Map{
				"success": false,
				"message": error,
			},
		})
	}
	user := user{Name: userDataReceived.Name, Email: userDataReceived.Email, Password: string(hashedPassword)}

	db.Create(&user)
	c.SendStatus(200)
	return c.JSON(fiber.Map{
		"statusCode": 201,
		"data": fiber.Map{
			"success": true,
			"value": fiber.Map{
				"name":  user.Name,
				"email": user.Email,
				"age":   user.Age,
			},
		},
	})
}

func LoginUser(c *fiber.Ctx) error {
	//TODO: Implement user login
	fmt.Println(string(c.Body()))
	return c.SendStatus(200)
}
