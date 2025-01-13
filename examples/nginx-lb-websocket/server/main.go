package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Define the upgrader with default options
var upgrader = websocket.Upgrader{
	// Allow all origins for simplicity. In production, you should set this to a specific origin.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	// Start the server on port 8080
	fmt.Println("WebSocket server starting on :8080/ws")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected")

	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error: %v", err)
			} else {
				log.Printf("Read error: %v", err)
			}
			break
		}

		// Print the received message
		fmt.Printf("Received: %s\n", message)

		// Prepare a response message
		response := fmt.Sprintf("Message received: %s", message)

		// Write response back to client
		err = conn.WriteMessage(messageType, []byte(response))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}

	log.Println("Client disconnected")
}
