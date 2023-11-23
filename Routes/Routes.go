package routes

import (
	"golang-fiber/Controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", controllers.HomeIndex)
	app.Get("/about", controllers.HomeAbout)
	app.Get("/hello/:name", controllers.HomeSayHello)

	// Mutliple handler dengan use (hanya untuk use method lain tidak bisa multiple handler) 
	// untuk multiple handler. handler yang merender (ApiIndex) harus dibelakang
	// jika tidak maka handler yang lainnya tidak akan diesekusi.
	app.Use("/api", controllers.ApiSetCustomHeader, controllers.ApiIndex)

	// route static digunakan untuk agar client bisa mengakses static kita
	app.Static("/static", "./Public/static")

}
