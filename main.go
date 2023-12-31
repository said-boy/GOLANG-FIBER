package main

import (
	// "crypto/tls"
	config "golang-fiber/Config"
	"golang-fiber/ErrorHandler"
	routes "golang-fiber/Routes"
	"log"

	// "net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2" // go get dulu, baru bisa pakai
	"github.com/qinains/fastergoding"
)

func main() {
	fastergoding.Run()
	engine := html.New("./Views", ".gohtml")

	app := fiber.New(fiber.Config{
		AppName:           "Aplikasi Web Golang",
		EnablePrintRoutes: true,
		Views: engine,
		// ViewsLayout: "./Views/main.gohtml",
		ErrorHandler: errorhandler.ErrorHandler, // mengganti default dengan custom
	})

	// menambahkan recover ke semua route, jadi semua panic akan dapat ditangani
	// dan tidak langsung menyebabkan server mati.
	app.Use(recover.New())

	config.LoadConfig(app)

	routes.SetupRoutes(app)

	// dapat membuka port di https
	// log.Fatal(app.ListenTLS(":3000", "./Cert/cert.pem", "./Cert/key.pem"))

	log.Fatal(app.Listen(":3000"))
}
