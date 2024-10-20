package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrade the HTTP connection to a WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: false, // Disable WebSocket compression
}

// WebSocket handler
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP request to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		// Read message from WebSocket
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		// Print received message
		fmt.Printf("Received: %s\n", message)

		// Echo the message back to the client
		if err := conn.WriteMessage(messageType, message); err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("my name is LOL")
	fmt.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
