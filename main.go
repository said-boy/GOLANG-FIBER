package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:           "Aplikasi Web Golang",
		EnablePrintRoutes: true,
	})
	app.Get("/:name/:last", func(c *fiber.Ctx) error {
		param := c.AllParams()
		for _, value := range param {
			fmt.Println(value)
		}
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
