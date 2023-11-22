package routes

import 	(
	"golang-fiber/Controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){

	app.Get("/", controllers.HomeIndex)
	app.Get("/about", controllers.HomeAbout)

}