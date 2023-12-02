package main

import (
	// "crypto/tls"
	config "golang-fiber/Config"
	routes "golang-fiber/Routes"
	"log"
	// "net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2" // go get dulu, baru bisa pakai
	"github.com/qinains/fastergoding"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	fastergoding.Run()
	engine := html.New("./Views", ".gohtml")

	app := fiber.New(fiber.Config{
		AppName:           "Aplikasi Web Golang",
		EnablePrintRoutes: true,
		Views: engine,
		ViewsLayout: "Views/main",
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
