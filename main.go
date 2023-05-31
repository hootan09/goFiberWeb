package main

import (
	"embed"
	"goV2Web/configs"
	"goV2Web/database"
	"goV2Web/middleware"
	api_routes "goV2Web/routes/api"
	web_routes "goV2Web/routes/web"
	"goV2Web/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "goV2Web/docs" // load API Docs files (Swagger)

	"github.com/gofiber/swagger" // swagger handler

	// jwtware "github.com/gofiber/jwt"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

//go:embed views/*
var viewsfs embed.FS

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization(Bearer)
// @host localhost:3000
// @BasePath /api
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
		cors.New(cors.Config{
			AllowCredentials: true,
		}),
		// Add simple logger.
		logger.New(),
	)

	setupRoutes(app)
	// app.Listen(os.Getenv("SERVER_URL"))
	utils.StartServer(app)
}

func setupRoutes(app *fiber.App) {

	app.Static("/", "./views/public")

	// JWT Middleware
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	// }))
	// Api routes
	api := app.Group("/api/v1")
	api.Get("/users", api_routes.GetAllUsers_api)
	api.Post("/user", api_routes.SaveUsers_api)
	api.Post("/token/new", api_routes.GetNewAccessToken)
	api.Get("/restricted", middleware.JWTProtected(), api_routes.Restricted_api)
	api.Post("/upload", middleware.JWTProtected(), api_routes.Upload_api)
	api.Get("/csv", middleware.JWTProtected(), api_routes.CSV_api)

	// Template Engine
	web := app.Group("/")
	web.Get("/", web_routes.Index_web)
	web.Get("/login", web_routes.Login_web)
	web.Post("/login", web_routes.Login_web_post)
	web.Get("/logout", web_routes.Logout_web)

	web.Get("/widgets.html", web_routes.Widgets_web)
	web.Get("/charts.html", web_routes.Charts_web)
	web.Get("/elements.html", web_routes.Elements_web)
	web.Get("/panels.html", web_routes.Panels_web)

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
			//for frontEnd
			// return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
		},
	)
}
