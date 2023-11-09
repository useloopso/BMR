package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// mainnet "github.com/useloopso/BMR/networks"
	config "github.com/useloopso/BMR/config"
	router "github.com/useloopso/BMR/router"
)

func main() {
	// load config
	config, err := config.LoadMainnetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// new fiber app
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format e.g. "localhost, loopso.xyz"
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Set-up routes
	router.Init(app)

	// Listen to DAI events
	// mainnet.ListenToEvents()

	// listen to port 3000
	log.Fatal(app.Listen(":" + config.PORT))
}
