package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	// WebSocket server URL
	url := "ws://localhost:8080/ws"
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()

	// Disable client-side compression
	conn.EnableWriteCompression(false)

	// Send a message to the server
	err = conn.WriteMessage(websocket.TextMessage, []byte("HELLO SERVER"))
	if err != nil {
		log.Fatal("Error sending message:", err)
	}

	// Read a message from the server
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("Read error:", err)
	}
	fmt.Printf("Received from server: %s\n", message)
}
