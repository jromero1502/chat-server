package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	incommingMessages    chan []byte
	incommingConnections chan *websocket.Conn
	connections          []*websocket.Conn
}

var upgrader = websocket.Upgrader{}

func NewWebSocket() *WebSocket {
	ws := &WebSocket{
		make(chan []byte),
		make(chan *websocket.Conn),
		make([]*websocket.Conn, 0),
	}
	return ws
}

func (ws *WebSocket) HandleMessage(currentConn *websocket.Conn) {
	for {
		select {
		case conn := <-ws.incommingConnections:
			fmt.Println("New connection")
			ws.connections = append(ws.connections, conn)
		case msg := <-ws.incommingMessages:
			PrintServerInfo("New message incomming. " + string(msg))
			currentConn.WriteMessage(websocket.TextMessage, []byte("He recibido tu mensaje :)"))
		}
	}
}

func (ws *WebSocket) Listen(w http.ResponseWriter, request *http.Request) {
	conn, _ := upgrader.Upgrade(w, request, nil)
	go ws.HandleMessage(conn)
	defer conn.Close()
	ws.incommingConnections <- conn
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			PrintServerInfo("Conexion de socket cerrada")
			return
		}
		ws.incommingMessages <- msg
	}
}
