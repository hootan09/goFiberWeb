package main

import (
	"embed"
	"goV2Web/configs"
	"goV2Web/database"
	api_routes "goV2Web/routes/api"
	web_routes "goV2Web/routes/web"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "goV2Web/docs" // load API Docs files (Swagger)

	"github.com/gofiber/swagger" // swagger handler

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

//go:embed views/*
var viewsfs embed.FS

func setupRoutes(app *fiber.App) {

	app.Static("/", "./views/public")

	// Api routes
	api := app.Group("/api/v1")
	api.Get("/users", api_routes.GetAllUsers_api)
	api.Post("/user", api_routes.SaveUsers_api)

	// Template Engine
	web := app.Group("/")
	web.Get("/", web_routes.Index_web)

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
	config := configs.FiberConfig(viewsfs)
	app := fiber.New(config)

	dbErr := database.InitDatabase()
	if dbErr != nil {
		panic(dbErr)
	}

	app.Use(
		// Add CORS to each route.
		cors.New(),
		// Add simple logger.
		logger.New(),
	)
	setupRoutes(app)
	app.Listen(os.Getenv("SERVER_URL"))
}
