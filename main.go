package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	handlers "github.com/useloopso/BMR/handlers"
	// mainnet "github.com/useloopso/BMR/networks"
	router "github.com/useloopso/BMR/router"
)

func main() {
	// new fiber app
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format e.g. "localhost, loopso.xyz"
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Set-up routes
	router.Initalize(app)

	// WebSocket route
	app.Get("/ws", func(c *fiber.Ctx) error {
		// Upgrade the request to a WebSocket connection
		return c.SendStatus(http.StatusSwitchingProtocols)
	})

	go handlers.HandleWebSocketConnections() // Start WebSocket handler

	// Listen to DAI events
	// mainnet.ListenToEvents()

	// listen to port 3000
	app.Listen(":3000")
}
