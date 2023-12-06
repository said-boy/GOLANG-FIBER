package controllers

import (
	"golang-fiber/Models/auth"
	"golang-fiber/Models/register"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()

func FormIndex(c *fiber.Ctx) error {
	return c.Render("form/index", fiber.Map{}) 
}

func FormAuth(c *fiber.Ctx) error {
	user := auth.Auth{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	err := Validate.Struct(user)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).
		JSON(fiber.Map{
			"Username" : "Username is required",
			"Password" : "Password is required",
		})
	}
	return c.SendStatus(200)
}

func FormRegister(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodGet {
		return c.Render("form/register", nil)
	}

	register := register.Register{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
		ConfirmPassword: c.FormValue("confirmpassword"),
	}

	err := Validate.StructExcept(register)
	if err != nil {
		validationError := err.(validator.ValidationErrors)
		for _, fieldError := range validationError {
			return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				fieldError.Field():fieldError.Tag(),
			})
		}
	}

	return c.SendStatus(200)
}