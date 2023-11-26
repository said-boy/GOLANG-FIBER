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

	// routing Group
	// untuk mengelompokkan pada 1 parent route

	v1 := app.Group("/v1") // v1/
	v1.Get("/", controllers.V1GroupIndex) // v1/
	v1.Get("/about", controllers.V1GroupAbout) // v1/about
	
	v2 := app.Group("/v2") // v2/
	v2.Get("/", controllers.V2GroupIndex) // v2/
	v2.Get("/about", controllers.V2GroupAbout) // v2/about

	// PR . menambahkan middleware ke group dan route
	app.Route("/route", func(router fiber.Router) {
		router.Get("/satu", controllers.RouteSatuIndex).Name("satu") // "route.satu"
		router.Get("/dua", controllers.RouteDuaIndex).Name("dua") // "route.dua"
	}, "route.")

	ctx := app.Group("/ctx")
	ctx.Get("/accept", controllers.CtxAccept).Name("accept") 

}
