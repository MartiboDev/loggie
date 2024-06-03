package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWSOrderbook(ws *websocket.Conn) {
	fmt.Println("New connection from client: ", ws.RemoteAddr())

	for {
		payload := fmt.Sprint("orderbook data -> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep((time.Second * 2))
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New connection from client: ", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Read error: ", err)
			continue
		}
		msg := buf[:n]

		fmt.Println("Chat: ", string(msg))

		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			_, err := ws.Write(b)
			if err != nil {
				fmt.Println("Write error: ", err)
			}
		}(ws)
	}
}

const Port int64 = 3000

func main() {
	fmt.Printf("Starting loggie on port %d\n", Port)

	server := NewServer()

	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/orderbookfeed", websocket.Handler(server.handleWSOrderbook))
	http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
}
