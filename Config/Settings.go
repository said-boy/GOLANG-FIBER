package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func LoadConfig(app *fiber.App){
	// MaxConnsPerIp untuk menghindari serangan DDOS
	// ini mengsetting berapa banyak koneksi yang dapat dilakukan oleh 1 ip.
	app.Server().MaxConnsPerIP = 1

	// Concurrency Mengatur berapa banyak orang yang dapat mengakses server secara bersamaan
	// jika batas sudah tercapai, maka client yang ingin mengakses akan disuruh menunggu hingga orang yang mengakses berhenti mengakses server
	app.Server().Concurrency = 1

	app.Server().CloseOnShutdown = true

	// ContinueHandler ketika client mengirimkan Expect 100-continue maka ini akan dipanggil
	// tentang header Expect 100-continue bisa browsing dichatgpt
	app.Server().ContinueHandler = func(header *fasthttp.RequestHeader) bool {
		// tambahkan logic disini
		return true
	}

	// secara default (false) header request/response akan diubah sesuai aturannya (normalnya).
	// aturannya (normalnya) adalah huruf depan setiap kata menjadi besar, spasi diganti - , huruf sisanya kecil
	// contoh: penulisan -> cONTENT tYpe; normalnya diubah menjadi Content-Type
	// jika ini diset true maka tidak akan diatur (normalkan), dan akan dikirim sesuai penulisan.
	app.Server().DisableHeaderNamesNormalizing = false

	// jika diset true maka server akan menutup koneksinya setelah membalas response untuk yang pertama kali.
	app.Server().DisableKeepalive = false

	// secara default (false) ketika ada form multipart/file akan diparsing oleh server
	// parsing sendiri tujuannya agar data bisa digunakan, jika menyetingnya ke true , 
	// maka server tidak akan memparsingnya, dan kita harus melakukan parsing manual pada handler/middleware saat kita membutuhkan datanya.
	app.Server().DisablePreParseMultipartForm = false

}