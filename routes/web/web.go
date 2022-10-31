package web

import "github.com/gofiber/fiber/v2"

func Index_web(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
		"Body":  "Hello mamad niki",
	})
}
