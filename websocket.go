package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	connections []*websocket.Conn
}

var upgrader = websocket.Upgrader{}

func NewWebSocket() *WebSocket {
	ws := &WebSocket{
		make([]*websocket.Conn, 0),
	}
	return ws
}

func (ws *WebSocket) HandleMessage(stateChannel chan bool, connChannel chan *websocket.Conn, msgChannel chan []byte, currentConn *websocket.Conn) {
	for {
		select {
		case conn := <-connChannel:
			fmt.Println("New connection")
			ws.connections = append(ws.connections, conn)
		case msg := <-msgChannel:
			PrintServerInfo("New message incomming. " + string(msg))
			currentConn.WriteMessage(websocket.TextMessage, []byte("He recibido tu mensaje :)"))
		case state := <-stateChannel:
			if !state {
				return
			}
		}
	}
}

func (ws *WebSocket) CheckOrigin(r *http.Request) bool {
	return true
}

func (ws *WebSocket) Listen(w http.ResponseWriter, request *http.Request) {
	upgrader.CheckOrigin = ws.CheckOrigin
	conn, err := upgrader.Upgrade(w, request, nil)
	if err != nil {
		PrintServerInfo("There was an error on the upgrader")
		PrintServerInfo(err.Error())
		return
	}
	connChannel := make(chan *websocket.Conn)
	msgChannel := make(chan []byte)
	stateChannel := make(chan bool)
	go ws.HandleMessage(stateChannel, connChannel, msgChannel, conn)
	defer conn.Close()
	connChannel <- conn
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			PrintServerInfo("Conexion de socket cerrada")
			stateChannel <- false
			return
		}
		msgChannel <- msg
	}
}
