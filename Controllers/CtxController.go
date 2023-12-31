package controllers

import (
	"fmt"

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

// akan mengambil semua parameter yang dikirimkan.
// untuk mengambilnya bisa menggunakan indexing atau looping
func CtxAllParams(c *fiber.Ctx) error {
	allParams := c.AllParams()
	fmt.Println(allParams["*1"]) // indexing

	//looping
	for key, value := range allParams {
		fmt.Println("key : "+key, "|| value : "+value)
	}
	return c.Next()
}

// stack mengembalikan [route, name, method, params] yang ada pada aplikasi
func CtxStack(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}

// menambahkan ke header response berupa {key} dan {value}
func CtxAppend(c *fiber.Ctx) error {
	c.Append("Link", "https://google.com")
	c.Append("Link-Dua", "https://facebook.com", "https://tiktok.com")
	return c.Next()
}

// kontent yang dikirimkan ke response tidak akan di render dibrowser
// melainkan akan otomatis terdownload oleh client
func CtxAttachment(c *fiber.Ctx) error {
	// "sayHello.png" adl nama file yang akan diberikan ke client saat file terdownload 
	// ekstensi file harus sama dengan exktensi yang akan dikirimkan oleh server,
	// jika tidak maka gambar tidak akan bisa dibuka. 
	// harus direname secara manual oleh client untuk dapat dibuka
	c.Attachment("sayHello.png") 

	// response yang akan dikirimkan ke client berupa file gambar dari path berikut
	// file dapat berupa text, gambar, file dan dokumen lainnya.
	return c.SendFile("./Public/static/image/civic.png") 
}