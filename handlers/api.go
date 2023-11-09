package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	mainnet "github.com/useloopso/BMR/networks"
)

var err *error

// template
func ApiHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "API request handled",
	})
}

func Home(c *fiber.Ctx) error {
	err := c.Status(200).JSON(fiber.Map{
		"success": true,
	})
	return err
}

func FetchBalance(c *fiber.Ctx) error {
	mainnet.FetchBalance(c)
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return nil
}

func FetchBlock(c *fiber.Ctx) error {
	mainnet.FetchBlock(c)
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return nil
}
