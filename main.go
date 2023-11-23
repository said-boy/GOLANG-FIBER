package main

import (
	"golang-fiber/Config"
	"golang-fiber/Routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:           "Aplikasi Web Golang",
		EnablePrintRoutes: true,
	})

	config.LoadConfig(app)

	routes.SetupRoutes(app)
	
	log.Fatal(app.Listen(":3000"))
}
