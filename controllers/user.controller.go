package controllers

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
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
	db.AutoMigrate(&User{})
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
			"success": false,
			"data": fiber.Map{
				"statusCode": 400,
				"message":    "Please provide all the details",
			},
		})
	}

	// if userDataReceived.

	// hashing the password
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(userDataReceived.Password), bcrypt.DefaultCost)
	if error != nil {
		c.SendStatus(500)
		return c.JSON(fiber.Map{
			"success": false,
			"data": fiber.Map{
				"statusCode": 500,
				"message":    error,
			},
		})
	}
	user := User{Name: userDataReceived.Name, Email: userDataReceived.Email, Password: string(hashedPassword)}

	db.Create(&user)
	c.SendStatus(201)
	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"statusCode": 201,
			"value": fiber.Map{
				"name":  user.Name,
				"email": user.Email,
				"age":   user.Age,
			},
		},
	})
}

func LoginUser(c *fiber.Ctx) error {

	// struct for getting the data from the user
	type userData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var userDataReceived userData
	if error := c.BodyParser(&userDataReceived); error != nil {
		log.Fatal(error)
	}

	// for getting data from database
	var userFound User
	error := db.Where("email = ?", userDataReceived.Email).First(&userFound).Error
	// if user not found then throw error
	if error != nil {
		if error == gorm.ErrRecordNotFound {
			// Record not found error handling
			c.SendStatus(401)
			return c.JSON(fiber.Map{
				"success": "false",
				"data": fiber.Map{
					"statusCode": 401,
					"message":    "Please provide correct credentials",
				},
			})
		}
	}

	// check whether password is correct or not
	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userDataReceived.Password)); err != nil {
		c.SendStatus(401)
		return c.JSON(fiber.Map{
			"success": "false",
			"data": fiber.Map{
				"statusCode": 401,
				"message":    "Please provide correct credentials",
			},
		})
	}

	c.SendStatus(200)
	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"statusCode": 200,
			"value": fiber.Map{
				"name":  userFound.Name,
				"email": userFound.Email,
				"age":   userFound.Age,
			},
		},
	})

}
