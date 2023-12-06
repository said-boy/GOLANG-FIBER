package routes

import (
	controllers "golang-fiber/Controllers"
	middleware "golang-fiber/Middleware"

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

	v1 := app.Group("/v1")                     // v1/
	v1.Get("/", controllers.V1GroupIndex)      // v1/
	v1.Get("/about", controllers.V1GroupAbout) // v1/about

	v2 := app.Group("/v2")                     // v2/
	v2.Get("/", controllers.V2GroupIndex)      // v2/
	v2.Get("/about", controllers.V2GroupAbout) // v2/about

	// PR . menambahkan middleware ke group dan route
	app.Route("/route", func(router fiber.Router) {
		router.Get("/satu", controllers.RouteSatuIndex).Name("satu") // "route.satu"
		router.Get("/dua", controllers.RouteDuaIndex).Name("dua")    // "route.dua"
	}, "route.")

	ctx := app.Group("/ctx")
	ctx.Get("/accept", controllers.CtxAccept).Name("accept")

	// /all-params/halo/halolagi/ke-dua/hai => *1 = halo/halolagi, *2 = hai
	// atau bisa mengguanakan /:parameter
	ctx.Get("/all-params/*/ke-dua/*", controllers.CtxAllParams)

	ctx.Get("/stack", controllers.CtxStack).Name("stack")
	ctx.Get("/append", controllers.CtxAppend)
	ctx.Get("/attachment", controllers.CtxAttachment)

	// TEMPLATE ROUTES
	template := app.Group("/templates")
	template.Get("/", controllers.TemplatesIndex)
	template.Get("/sapa-:name?", controllers.TemplatesIndex)

	// jika ada route ke tempat yang sama, maka route yang akan digunakan/terpanggil
	// adalah route yang pertama kali ditulis.

	// http://localhost:3000/templates/@said
	// :sign => @ , :name => said
	// http://localhost:3000/templates/said
	// :sign => s , :name => aid
	// sign mengambil 1 karakter data parameter paling awal
	template.Get("/:sign:name", controllers.TemplatesSignRoutes)          // ini yang akan dipanggil
	template.Get("/:username<int>", controllers.TemplatesConstrainsRoute) // ini tidak terpanggil

	// memberi aturan pada parameter yang dikirimkan pada url.
	// http://localhost:3000/templates/username/123
	template.Get("/username/:username<int>", controllers.TemplatesConstrainsRoute)

	// http://localhost:3000/templates/username2/aaaaaaaZAaa // benar
	// http://localhost:3000/templates/username2/aaaaaaa123 // salah
	// alpha hanya memperbolehkan huruf a-z besar ataupun kecil
	template.Get("/username2/:username<alpha;minLen(8)>", controllers.TemplatesConstrainsRoute)

	// PENERAPAN MIDDLEWARE

	// karena middleware diletakkana di use
	// maka siapapun yang ingin mengakses /middleware harus memenuhi syaratnya.
	app.Use("/middleware", middleware.MiddlewareCustomCookie)

	midv1 := app.Group("/middleware")
	midv1.Get("/v1", controllers.V1GroupIndex)
	midv1.Get("/v1/about", controllers.V1GroupAbout)

	midv2 := app.Group("/middleware")
	midv2.Get("/v2", controllers.V2GroupIndex)
	midv2.Get("/v2/about", controllers.V2GroupAbout)

	// error handler
	app.Get("/error", controllers.ErrorHandler)

	// simulasi panic
	// by default, fiber tidak bisa menghandle panic. contoh..
	app.Get("/panic", controllers.ErrorPanicSimulation)
	// ketika diakses, maka server akan berhenti

	// cara menangani panic dengan menggunakan recover
	// dan meletakkanya di main.go

	// form
	form := app.Group("form")
	form.Get("/", controllers.FormIndex)
	form.Post("/auth", controllers.FormAuth)
	form.Get("/register", controllers.FormRegister)
	form.Post("/register", controllers.FormRegister)

}
