package main

import (
	"log"
	"golang-fiber/Routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:           "Aplikasi Web Golang",
		EnablePrintRoutes: true,
	})

	routes.SetupRoutes(app)
	
	log.Fatal(app.Listen(":3000"))
}
