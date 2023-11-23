package controllers

import "github.com/gofiber/fiber/v2"

// v1
func V1GroupIndex(c *fiber.Ctx) error {
	return c.SendString("Ini halaman Group Index v1")
}

func V1GroupAbout(c *fiber.Ctx) error {
	return c.SendString("Ini halaman Group About v1")
}

// v2
func V2GroupIndex(c *fiber.Ctx) error {
	return c.SendString("Ini halaman Group Index v2")
}

func V2GroupAbout(c *fiber.Ctx) error {
	return c.SendString("Ini halaman Group About v2")
}