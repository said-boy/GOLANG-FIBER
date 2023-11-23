package controllers

import "github.com/gofiber/fiber/v2"

func ApiIndex(c *fiber.Ctx) error {
	return c.SendString("Api")
}

func ApiSetCustomHeader(c *fiber.Ctx) error {
	c.Set("X-Custom-Header", "Halo")
	return c.Next()
}
