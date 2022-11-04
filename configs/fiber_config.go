package configs

import (
	"embed"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
)

// This stores all of your app's sessions
// Default middleware config
var Store *session.Store = session.New(session.Config{
	CookieSecure:   true,
	CookieHTTPOnly: true,
})

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig(viewFs embed.FS) fiber.Config {

	//cookie session Expiration
	expire, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	Store.Config.Expiration = time.Duration(expire) * time.Minute

	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	engine := html.NewFileSystem(http.FS(viewFs), ".html")
	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false

	// Debug will print each template that is parsed, good for debugging
	// engine.Debug(true) // Optional. Default: false

	// Layout defines the variable name that is used to yield templates within layouts
	// engine.Layout("embed") // Optional. Default: "embed"

	// Delims sets the action delimiters to the specified strings
	engine.Delims("{{", "}}") // Optional. Default: engine delimiters

	// AddFunc adds a function to the template's global function map.
	engine.AddFunc("greet", func(name string) string {
		return "Hello, " + name + "!"
	})

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		Views:       engine,
	}
}
