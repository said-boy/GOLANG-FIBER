package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// accept adalah fungdsi untuk bertanya kepada client, Content-Type ada yang dia bisa terima.
// Accept untuk memastikan kontent yang diberikan dari server dapat diterima oleh client

// server : kamu bisa menerima kontent bertipe json nggak? 
// client : bisa.
// server : yaudah, ini saya kirimkan kontent dalam bentuk json.

func CtxAccept(c *fiber.Ctx) error {
	accept := c.Accepts("application/json","plain/text")
	switch accept {

	case "application/json":
		return c.JSON(fiber.Map{
			"text":"Hallo",
			"status":fiber.StatusAccepted,
		})

	case "plain/text":
		return c.SendString("Halaman Accept")

	default:
		return c.Status(fiber.StatusNotAcceptable).SendString("Tidak diizinkan")
	
	}
}