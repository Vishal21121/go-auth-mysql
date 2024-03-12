package controllers

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
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

type (
	UserData struct {
		Name     string `json:"name" validate:"required,min=3,max=20"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	LoginData struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	ErrorResponse struct {
		Error        bool
		FailedField  string
		Tag          string
		Value        interface{}
		ErrorMessage string
	}

	XValidator struct {
		validator *validator.Validate
	}
)

var validate = validator.New()

func (v XValidator) Validate(data interface{}) []string {
	var validationErrors []string

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem string

			switch err.Tag() {
			case "required":
				elem = fmt.Sprintf("%s is required", err.Field())
			case "email":
				elem = fmt.Sprintf("%s is not a valid email", err.Field())
			case "min":
				elem = fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
			case "max":
				elem = fmt.Sprintf("%s must be at most %s characters long", err.Field(), err.Param())
			}
			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

var db *gorm.DB

func DBSetter(database *gorm.DB) {
	db = database
	db.AutoMigrate(&User{})
}

func RegisterUser(c *fiber.Ctx) error {

	myValidator := &XValidator{
		validator: validate,
	}

	// converting the recived data in the following struct
	var userDataReceived UserData
	if err := c.BodyParser(&userDataReceived); err != nil {
		log.Println(err)
	}

	if errs := myValidator.Validate(userDataReceived); len(errs) > 0 {
		c.SendStatus(422)
		return c.JSON(fiber.Map{
			"success": false,
			"data": fiber.Map{
				"statusCode": 422,
				"value":      errs,
			},
		})
	}

	// if userDataReceived.

	// check whether we already have user with the provided email id
	var userFound User
	db.Where("email = ?", userDataReceived.Email).First(&userFound)

	if userFound.Email != "" {
		c.SendStatus(400)
		return c.JSON(fiber.Map{
			"suceess": false,
			"data": fiber.Map{
				"statusCode": 400,
				"message":    "Please enter another email id",
			},
		})
	}

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

	myValidator := &XValidator{
		validator: validate,
	}

	var userDataReceived LoginData
	if error := c.BodyParser(&userDataReceived); error != nil {
		log.Fatal(error)
	}

	if errs := myValidator.Validate(userDataReceived); len(errs) > 0 {
		c.SendStatus(422)
		return c.JSON(fiber.Map{
			"success": false,
			"data": fiber.Map{
				"statusCode": 422,
				"value":      errs,
			},
		})
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
