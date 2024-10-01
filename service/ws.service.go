package service

import (
	"codecloud/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	EventPing = "ping"
)

type WSService struct {
	clents map[*websocket.Conn]bool
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWSService() *WSService {
	return &WSService{
		clents: make(map[*websocket.Conn]bool),
	}
}

func (wsService *WSService) ConnectWSAndListen(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	wsService.clents[conn] = true
	go wsService.listenForEvents(conn) // Start listening for events in a new goroutine
	log.Println("Client connected: ", conn.RemoteAddr())
}

func (wsService *WSService) listenForEvents(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message: ", err)
			delete(wsService.clents, conn)
			break
		}

		var event model.WSEvent
		err = json.Unmarshal(message, &event)
		if err != nil {
			log.Println("Invalid event format: ", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid event format"))
			continue
		}

		log.Println("Event received: ", event.Data)

		switch event.EventType {
		case EventPing:
			conn.WriteMessage(websocket.TextMessage, []byte("Pong"))
		default:
			conn.WriteMessage(websocket.TextMessage, []byte("Unknown event type"))
		}
	}
}
