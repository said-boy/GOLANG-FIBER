package controllers

import "github.com/gofiber/fiber/v2"

func ErrorHandler(c *fiber.Ctx) error {
	return c.SendFile("./Views/errors/error.gohtml")
}

func ErrorPanicSimulation(c *fiber.Ctx) error {
	panic("Simulasi panic..")
}