package controllers

import "github.com/gofiber/fiber/v2"

func RouteSatuIndex(c *fiber.Ctx) error {
	return c.SendString("Route Satu")
}

func RouteDuaIndex(c *fiber.Ctx) error {
	return c.SendString("Route Dua")
}