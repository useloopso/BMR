package handlers

import (
    "github.com/gorilla/websocket"
    // "github.com/gofiber/fiber/v2"
    "net/http"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func HandleWebSocketConnections() {
    // Handle WebSocket connections in a goroutine
    http.HandleFunc("/ws", handleWebSocket)
    http.ListenAndServe(":8080", nil)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    // Upgrade the request to a WebSocket connection
    conn, _ := upgrader.Upgrade(w, r, nil)
    defer conn.Close()

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            return
        }
        if messageType == websocket.TextMessage {
            // Handle WebSocket messages as needed
        }
    }
}
