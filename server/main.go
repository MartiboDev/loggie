package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
)

var logs []LoggieLog

type LoggieLog struct {
	ID        int64     `json:"id"`
	Source    int64     `json:"source"`
	Severity  string    `json:"severity"`
	Category  string    `json:"category"`
	Resource  string    `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

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

func logEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// encode the logs array into JSON and write it to the response
		logsJSON, err := json.Marshal(logs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(logsJSON)
	} else if r.Method == "POST" {
		// decode the request body into a new log
		var log LoggieLog
		err := json.NewDecoder(r.Body).Decode(&log)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.ID = int64(len(logs) + 1)
		log.Timestamp = time.Now()

		logs = append(logs, log)
		logBytes, _ := json.Marshal(log)
		websocketServer.broadcast(logBytes) // Convert log to []byte before passing it to websocketServer.broadcast()
	} else {
		// return a 405 Method Not Allowed if the request method is not GET or POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected")

	websocketServer.handleConn(ws)
}

func setupRoutes() {
	http.HandleFunc("/", logEndpoint)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Starting server on port:", 8080)

	setupRoutes()

	corsOrigins := handlers.AllowedOrigins([]string{"http://localhost:5173"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsOrigins, corsMethods, corsHeaders)(http.DefaultServeMux)))
}
