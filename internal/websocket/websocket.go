package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

type WebsocketServer struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *WebsocketServer {
	return &WebsocketServer{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *WebsocketServer) HandleConn(conn *websocket.Conn) {
	s.conns[conn] = true
}

func (s *WebsocketServer) Broadcast(b []byte) {
	for conn := range s.conns {
		if err := conn.WriteMessage(websocket.TextMessage, b); err != nil {
			log.Println(err)
			conn.Close()
			delete(s.conns, conn)
		}
	}
}
