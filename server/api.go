package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
)

const Port int64 = 8080
const frontendPort int64 = 5173

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

func SetupRoutes() {
	http.HandleFunc("/", logEndpoint)
	http.HandleFunc("/ws", wsEndpoint)
}

func SetupCors() *http.Handler {
	frontendAddr := fmt.Sprint("http://localhost:", frontendPort)

	origins := handlers.AllowedOrigins([]string{frontendAddr})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	headers := handlers.AllowedHeaders([]string{"Content-Type"})

	handler := handlers.CORS(origins, methods, headers)(http.DefaultServeMux)

	return &handler
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
