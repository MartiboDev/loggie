package router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func logEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetLogs(w, r)
	case "POST":
		handlePostLog(w, r)
	default:
		// return a 405 Method Not Allowed if the request method is not GET or POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func handleGetLogs(w http.ResponseWriter, _ *http.Request) {
	// encode the logs array into JSON and write it to the response
	logsJSON, err := json.Marshal(logs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(logsJSON)
}

func handlePostLog(w http.ResponseWriter, r *http.Request) {
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
	websocketServer.Broadcast(logBytes) // Convert log to []byte before passing it to websocketServer.broadcast()
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected")

	websocketServer.HandleConn(ws)
}
