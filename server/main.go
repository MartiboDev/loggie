package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/net/websocket"
)

const PORT = 3000

var websocketServer *WebsocketServer

// WebsocketServer represents a websocket server
// that can handle multiple connections
type WebsocketServer struct {
	conns map[*websocket.Conn]bool
}

// create a new websocket server
func NewServer() *WebsocketServer {
	return &WebsocketServer{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *WebsocketServer) handleConn(ws *websocket.Conn) {
	s.conns[ws] = true
}

func (s *WebsocketServer) broadcast(b []byte) {
	for ws := range s.conns {
		fmt.Println("Broadcasting to: ", ws.RemoteAddr())
		if _, err := ws.Write(b); err != nil {
			fmt.Println("Write error: ", err)
			ws.Close()
			delete(s.conns, ws)
		}
	}
}

type LoggieLog struct {
	ID        int64     `json:"id"`
	Severity  string    `json:"severity"`
	Category  string    `json:"category"`
	Resource  string    `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

// create an array of logs to store log messages
var logs []LoggieLog

func logHandler(w http.ResponseWriter, r *http.Request) {
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
		var newLog LoggieLog
		err := json.NewDecoder(r.Body).Decode(&newLog)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log := createLog(newLog)
		logs = append(logs, log)
		logBytes, _ := json.Marshal(log)
		websocketServer.broadcast(logBytes) // Convert log to []byte before passing it to websocketServer.broadcast()
	} else {
		// return a 405 Method Not Allowed if the request method is not GET or POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func createLog(newLog LoggieLog) LoggieLog {
	return LoggieLog{
		ID:        int64(len(logs) + 1),
		Severity:  newLog.Severity,
		Category:  newLog.Category,
		Resource:  newLog.Resource,
		Timestamp: time.Now(),
		Message:   newLog.Message,
	}
}

func main() {
	fmt.Printf("Starting loggie on port %d\n", PORT)

	websocketServer = NewServer()

	r := mux.NewRouter()
	r.HandleFunc("/log", logHandler).Methods("GET", "POST")
	r.Handle("/ws", websocket.Handler(websocketServer.handleConn))

	handler := cors.Default().Handler(r)

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), handler)
}
