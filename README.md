
# WebSocket Chat Application

A simple WebSocket-based chat application built with Go using the `gorilla/websocket` package. This application demonstrates how to set up a WebSocket server in Go, handle multiple client connections, and broadcast messages to all connected clients.

## Features

- WebSocket-based chat communication
- Broadcasts messages to all connected clients
- Automatic reconnection handling on client-side (for the HTML frontend)
- Connection cleanup on server-side when clients disconnect

## Requirements

- Go 1.18+ (or newer)
- `gorilla/websocket` package (installed via `go get`)

## Running the Application

### 1. Clone the repository:

```bash
git clone https://github.com/sriniwas14/GolangWebsocketTutorial.git
cd GolangWebsocketTutorial
```

### 2. Install Go dependencies:

```bash
go mod tidy
```

### 3. Start the Go server:

```bash
go run main.go
```

The server will start running on `http://localhost:8080`.

### 4. Open the chat frontend:

- Open the `index.html` file in your browser.
- You can open multiple browser tabs or separate browsers to simulate multiple users.
- Messages typed in one client will be broadcasted to all other connected clients.

## File Structure

```
.
├── main.go              # Go WebSocket server handling connections and messages
├── templates
│   └── index.html       # HTML frontend for the chat app
```

## How It Works

- **Server**: The Go server listens for WebSocket connections on `/chat`. It maintains an active list of connections and broadcasts messages received from one client to all connected clients.
- **Client**: The frontend (`index.html`) connects to the server and allows users to send and receive chat messages. Messages are sent as JSON objects containing the username and the message text.

## Handling Disconnections

- When a client disconnects, the server cleans up the connection and removes it from the list of active connections.
- If a message fails to send to a client (e.g., due to the client disconnecting), the server will close the connection and remove it from the active list.

