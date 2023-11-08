package main

import (
	"github.com/gofiber/fiber/v2"
	mainnet "github.com/useloopso/BMR/networks"
)

func main() {
	// new fiber app
	app := fiber.New()

	// Fetch Balance
	mainnet.FetchBalance()

	// send a get request of type String
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Hello World!")
		return err
	})

	// Listen to DAI events
	mainnet.ListenToEvents()

	// listen to port 3000
	app.Listen(":3000")
}
