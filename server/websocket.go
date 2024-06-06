package main

import (
	"log"

	"github.com/gorilla/websocket"
)

var websocketServer = NewServer()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebsocketServer struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *WebsocketServer {
	return &WebsocketServer{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *WebsocketServer) handleConn(conn *websocket.Conn) {
	s.conns[conn] = true
}

func (s *WebsocketServer) broadcast(b []byte) {
	for conn := range s.conns {
		if err := conn.WriteMessage(websocket.TextMessage, b); err != nil {
			log.Println(err)
			conn.Close()
			delete(s.conns, conn)
		}
	}
}
