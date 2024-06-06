package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

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
		websocketServer.Broadcast(logBytes) // Convert log to []byte before passing it to websocketServer.broadcast()
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

	websocketServer.HandleConn(ws)
}
