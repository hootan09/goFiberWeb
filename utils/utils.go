package utils

import (
	"goV2Web/configs"

	"github.com/gofiber/fiber/v2"
)

func Authorized_web(c *fiber.Ctx) bool {
	sess, err := configs.Store.Get(c)
	if err != nil {
		panic(err)
	}
	authorized := sess.Get("authorized")
	return (authorized != nil && authorized == "true")
}
