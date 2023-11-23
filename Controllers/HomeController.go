package controllers

import "github.com/gofiber/fiber/v2"

func HomeIndex(c *fiber.Ctx) error {
	return c.SendString("Halaman Home")
}

func HomeAbout(c *fiber.Ctx) error {
	return c.SendString("Halaman About")
}

func HomeSayHello(c *fiber.Ctx) error {
	return c.SendString("Hello " + c.Params("name", "Guest"))
}
