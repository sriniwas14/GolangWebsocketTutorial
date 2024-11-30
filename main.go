package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Msg struct {
	Username string `json:"username"`
	Msg      string `json:"msg"`
}

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Msg)
	upgrader  = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
)

// Serve the HTML template
func serveHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Handle WebSocket connections
func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	clients[conn] = true
	defer func() {
		delete(clients, conn)
		conn.Close()
	}()
	for {
		var msg Msg
		if err := conn.ReadJSON(&msg); err != nil {
			log.Print("Error reading message", err)
			break
		}
		broadcast <- msg
	}
}

// Broadcast messages to all clients
func handleMessages() {
	for msg := range broadcast {
		log.Print(msg)
		for client := range clients {
			_ = client.WriteJSON(msg)
		}
	}
}

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/chat", handleConnection)
	go handleMessages()

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
