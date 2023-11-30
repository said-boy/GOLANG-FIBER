package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func TemplatesIndex(c *fiber.Ctx) error {
	text := c.Params("name", "di Index Template")

	return c.Render("index", fiber.Map{
		"title": "Index Template",
		"welcome": "Selamat datang "+text,
	})

}