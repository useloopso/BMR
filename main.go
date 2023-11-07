package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// new fiber app
	app := fiber.New()

	// send a get request of type String
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Hello World!")
		return err
	})

	// listen to port 3000
	app.Listen(":3000")
}
