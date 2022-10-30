package main

import (
	"goV2Web/configs"
	"goV2Web/database"
	"goV2Web/routes"
	"os"

	"github.com/gofiber/fiber/v2"

	_ "goV2Web/docs" // load API Docs files (Swagger)

	"github.com/gofiber/swagger" // swagger handler

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App) {

	// Template Engine must goes here
	app.Get("/", status)

	// Api routes
	api := app.Group("/api/v1")
	api.Get("/users", routes.GetAllUsers)
	api.Post("/user", routes.SaveUsers)

	// Swagger routes
	swag := app.Group("/apidoc")
	// Default
	swag.Get("*", swagger.HandlerDefault)
	// custom
	// swag.Get("/swagger/*", swagger.New(swagger.Config{
	// 	URL: "http://example.com/doc.json",
	// 	DeepLinking: false,
	// 	// Expand ("list") or Collapse ("none") tag groups by default
	// 	DocExpansion: "none",
	// 	// Prefill OAuth ClientId on Authorize popup
	// 	OAuth: &swagger.OAuthConfig{
	// 		AppName:  "OAuth Provider",
	// 		ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
	// 	},
	// 	// Ability to change OAuth2 redirect uri location
	// 	OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	// }))

	// Register new special NotFound route.
	app.Use(
		// Anonymous function.
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "sorry, endpoint is not found",
				"data":    nil,
			})
		},
	)
}

func main() {

	// Define Fiber Config
	config := configs.FiberConfig()

	app := fiber.New(config)

	dbErr := database.InitDatabase()
	if dbErr != nil {
		panic(dbErr)
	}

	setupRoutes(app)
	app.Listen(os.Getenv("SERVER_URL"))
}
