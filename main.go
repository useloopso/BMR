package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	config "github.com/useloopso/BMR/config"
	router "github.com/useloopso/BMR/router"
)

func main() {
	config, err := config.LoadMainnetConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format e.g. "localhost, loopso.xyz"
		AllowHeaders: "Origin, Content-Type, Acceptlocalhost",
	}))

	// Set-up routes
	router.Init(app)

	// listen to port 3000
	log.Fatal(app.Listen("localhost:" + config.PORT))
}
