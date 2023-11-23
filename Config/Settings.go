package config

import "github.com/gofiber/fiber/v2"

func LoadConfig(app *fiber.App){
	// MaxConnsPerIp untuk menghindari serangan DDOS
	// ini mengsetting berapa banyak koneksi yang dapat dilakukan oleh 1 ip.
	app.Server().MaxConnsPerIP = 1
}