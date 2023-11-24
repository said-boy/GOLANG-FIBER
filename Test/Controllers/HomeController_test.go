package controllers_test

import (
	"golang-fiber/Routes"
	"io"
	"net/http/httptest"
	"testing"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var app = fiber.New()

// fungsi init akan selalu diesekusi diawal setiap kali fungsi yang lain (yaitu test dalam file ini) akan dijalankan.
func init(){
	routes.SetupRoutes(app)
}

func TestHomeIndex(t *testing.T) {
	// membuat request
	req := httptest.NewRequest(fiber.MethodGet, "http://localhost:3000/", nil)

	resp, err := app.Test(req)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 200, resp.StatusCode)
	
	// mengambil string body response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "Halaman Home", string(body))

}