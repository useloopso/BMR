package router

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/useloopso/BMR/handlers"
)

func Initalize(app *fiber.App) {
	// Home
	app.Get("/", handlers.Home)

	// Fetch Balance
	app.Get("/balance", handlers.FetchBalance)

	// API route
	app.Get("/api", handlers.ApiHandler)
}
