package controllers

import (
	"golang-fiber/Models/auth"
	"github.com/gofiber/fiber/v2"
)

func FormIndex(c *fiber.Ctx) error {
	return c.Render("form/index", fiber.Map{}) 
}

func FormAuth(c *fiber.Ctx) error {
	user := auth.User{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	err := auth.Validate.Struct(user)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).
		JSON(fiber.Map{
			"Username" : "Username is required",
			"Password" : "Password is required",
		})
	}
	return c.SendStatus(200)
}