package api

import (
	"encoding/csv"
	"fmt"
	"goV2Web/database"
	"goV2Web/models"
	"goV2Web/utils"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// GetNewAccessToken method for create a new access token.
// @Description Create a new access token.
// @Summary create a new access token
// @Tags Token
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} status "ok"
// @Router /api/v1/token/new [post]
func GetNewAccessToken(c *fiber.Ctx) error {
	admin := &models.Admin{}
	if err := c.BodyParser(admin); err != nil {
		return err
	}
	if admin.Username != os.Getenv("ADMIN_USERNAME") || admin.Password != os.Getenv("ADMIN_PASSWORD") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "invalid username password!",
		})
	}

	// Generate a new Access token.
	token, err := utils.GenerateNewAccessToken(admin)
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sucess":  true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success":      true,
		"message":      nil,
		"access_token": token,
	})
}

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
func SaveUsers_api(c *fiber.Ctx) error {
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
func GetAllUsers_api(c *fiber.Ctx) error {
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

// Restricted func get username from token
// @Description Get username from token.
// @Summary get username from token
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Security ApiKeyAuth
// @Router /api/v1/restricted [get]
func Restricted_api(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return c.SendString("Welcome " + username)
}

// upload func upload single file to server
// @Description upload single file to server.
// @Summary upload single file to server
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Security ApiKeyAuth
// @Router /api/v1/upload [post]
func Upload_api(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["file"]
	filename := ""

	for _, file := range files {
		filename = file.Filename
		fmt.Println(filename)
		if err := c.SaveFile(file, "./views/public/uploads/"+filename); err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{
		"url": "./uploads/" + filename,
	})
}

// csv func Download csv
// @Description Download csv
// @Summary Download csv
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Security ApiKeyAuth
// @Router /api/v1/upload [post]
func CSV_api(c *fiber.Ctx) error {
	filePath := "./views/public/uploads/users.csv"

	if err := CreateFile(filePath); err != nil {
		return err
	}

	return c.Download(filePath)
}

func CreateFile(filePath string) error {
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	users, err := database.GetAllUsers()
	if err != nil {
		return err
	}

	writer.Write([]string{
		"ID", "Email", "Active", "CreatedAt", "UpdatedAt",
	})

	for _, user := range users {
		data := []string{
			strconv.Itoa(int(user.ID)),
			user.Email,
			strconv.FormatBool(user.Active),
			"",
			user.UpdatedAt.String(),
		}

		if err := writer.Write(data); err != nil {
			return err
		}
	}

	return nil
}
