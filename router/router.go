package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	handlers "github.com/useloopso/BMR/handlers"
	middleware "github.com/useloopso/BMR/middleware"
)

func Init(router *fiber.App) {
	router.Use(middleware.Security)
	router.Use(middleware.Json)

	router.Get("/", handlers.Home)

	v1 := router.Group("/v1")
	v1.Get("/balance", handlers.FetchBalance)
	v1.Get("/block", handlers.FetchBlock)
	v1.Get("/api", handlers.ApiHandler)

	ws := v1.Group("/ws", middleware.UpgradeWebsocketConnection)
	ws.Get("/:id", websocket.New(handlers.WebsocketHandler))

	// Error handling
	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "404: Not Found",
		})
	})
}
