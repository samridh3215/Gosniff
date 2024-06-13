package main

import (
	"log"
	"net/http"
	"sync"
	"github.com/gorilla/websocket"
	// "nhooyr.io/websocket"
)

var (
	connections = make([]*websocket.Conn, 0)
	connMutex   = &sync.Mutex{}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	connMutex.Lock()
	connections = append(connections, ws)
	connMutex.Unlock()
	infoLog.Println("New client connected")
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			infoLog.Println("Client disconnected")
			break
		}
	}

}
func removeConnection(conns []*websocket.Conn, conn *websocket.Conn) []*websocket.Conn {
	index := -1
	for i, c := range conns {
		if c == conn {
			index = i
			break
		}
	}
	if index != -1 {
		conns = append(conns[:index], conns[index+1:]...)
	}
	return conns
}



func BroadcastToConnections(message string) {
	connMutex.Lock()
	defer connMutex.Unlock()
	for _, conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			panicLog.Println("Error sending message:", err)
			// Remove the connection if there's an error
			conn.Close()
			connections = removeConnection(connections, conn)
		}
	}
}
