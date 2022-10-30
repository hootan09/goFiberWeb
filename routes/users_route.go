package routes

import (
	"goV2Web/database"
	"goV2Web/models"

	"github.com/gofiber/fiber/v2"
)

// SaveUsers func for creates a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param active body boolean true "Active"
// @Success 200 {object} models.Users
// @Security ApiKeyAuth
// @Router /api/v1/user [post]
func SaveUsers(c *fiber.Ctx) error {
	newUser := new(models.Users)

	err := c.BodyParser(newUser)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	result, err := database.CreateUser(newUser.Email, newUser.Active)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
	return nil
}

// GetAllUsers func gets all exists users.
// @Description Get all exists users.
// @Summary get all exists users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.Users
// @Router /api/v1/users [get]
func GetAllUsers(c *fiber.Ctx) error {
	result, err := database.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}
