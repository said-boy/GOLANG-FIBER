package middleware

import "github.com/gofiber/fiber/v2"

func MiddlewareCustomCookie(c *fiber.Ctx) error {
	cookie := c.Cookies("key")
	if cookie == "" {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Next()
}