package web

import (
	"goV2Web/configs"
	"goV2Web/models"
	"goV2Web/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Index_web(c *fiber.Ctx) error {

	if !utils.Authorized_web(c) {
		return c.Redirect("/login")
	} else {
		return c.Render("views/index", fiber.Map{
			"Title": "Dashboard",
		}, "views/layouts/main")
	}
}

func Login_web(c *fiber.Ctx) error {
	if !utils.Authorized_web(c) {
		return c.Render("views/login", fiber.Map{})
	}
	return c.Redirect("/")
}

func Login_web_post(c *fiber.Ctx) error {
	admin := &models.Admin{}
	if err := c.BodyParser(admin); err != nil {
		return err
	}
	if admin.Username == os.Getenv("ADMIN_USERNAME") && admin.Password == os.Getenv("ADMIN_PASSWORD") {
		sess, err := configs.Store.Get(c)
		if err != nil {
			panic(err)
		}

		sess.Set("authorized", "true")
		if err := sess.Save(); err != nil {
			panic(err)
		}
	}

	return c.Redirect("/")

}

func Logout_web(c *fiber.Ctx) error {
	sess, err := configs.Store.Get(c)
	if err != nil {
		panic(err)
	}

	sess.Delete("authorized")
	if err := sess.Save(); err != nil {
		panic(err)
	}
	return c.Redirect("/")
}

// for test ui
func Charts_web(c *fiber.Ctx) error {
	if !utils.Authorized_web(c) {
		return c.Redirect("/login")
	} else {
		return c.Render("views/charts", fiber.Map{})
	}
}

func Elements_web(c *fiber.Ctx) error {
	if !utils.Authorized_web(c) {
		return c.Redirect("/login")
	} else {
		return c.Render("views/elements", fiber.Map{})
	}
}

func Panels_web(c *fiber.Ctx) error {
	if !utils.Authorized_web(c) {
		return c.Redirect("/login")
	} else {
		return c.Render("views/panels", fiber.Map{})
	}
}

func Widgets_web(c *fiber.Ctx) error {
	if !utils.Authorized_web(c) {
		return c.Redirect("/login")
	} else {
		return c.Render("views/widgets", fiber.Map{})
	}
}
