package handlers

import (
    "github.com/gofiber/fiber/v2"
	mainnet "github.com/useloopso/BMR/networks"
)

func Home(c *fiber.Ctx) error {
	response := []byte(`{"success": true}`)
	err := c.Status(200).Send(response)
	return err
}

func ApiHandler(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "API request handled"})
}

func FetchBalance(c *fiber.Ctx) error {
    mainnet.FetchBalance(c)
    return nil
}