package web

import "github.com/gofiber/fiber/v2"

var isLoggedIn = false

func Index_web(c *fiber.Ctx) error {
	// return c.Render("views/index", fiber.Map{
	// 	"Title": "Hello, World!",
	// 	"Body":  "Hello mamad niki",
	// }, "views/layouts/main")
	if !isLoggedIn {
		return c.Redirect("/login")
	} else {
		return c.Render("views/index", fiber.Map{
			"Title": "Hello, World!",
			"Body":  "Hello mamad niki",
		})
	}
}

func Login_web(c *fiber.Ctx) error {
	if !isLoggedIn {
		return c.Render("views/login", fiber.Map{})
	}
	return c.Redirect("/")
}

func Login_web_post(c *fiber.Ctx) error {
	isLoggedIn = true
	return c.Redirect("/")
}

func Logout_web(c *fiber.Ctx) error {
	isLoggedIn = false
	return c.Redirect("/")
}
