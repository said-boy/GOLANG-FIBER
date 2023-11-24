package main

import (
	// "crypto/tls"
	config "golang-fiber/Config"
	routes "golang-fiber/Routes"
	"log"
	// "net"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:           "Aplikasi Web Golang",
		EnablePrintRoutes: true,
	})

	config.LoadConfig(app)

	routes.SetupRoutes(app)

	// dapat membuka port di https
	// log.Fatal(app.ListenTLS(":3000", "./Cert/cert.pem", "./Cert/key.pem"))

	log.Fatal(app.Listen(":3000"))
}
