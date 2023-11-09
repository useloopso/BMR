// In handlers/websocket.go
package handlers

import (
	"log"

	"github.com/gofiber/websocket/v2"

	mainnet "github.com/useloopso/BMR/networks"
)

func WebsocketHandler(c *websocket.Conn) {
	// c.Locals is added to the *websocket.Conn
	log.Println(c.Locals("allowed"))  // true
	log.Println(c.Params("id"))       // 123
	log.Println(c.Query("v"))         // 1.0
	log.Println(c.Cookies("session")) // ""

	// Listen to DAI approval events
	mainnet.ListenToEvents(c)

}
