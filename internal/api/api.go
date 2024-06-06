package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	wsHelper "github.com/MartiboDev/loggie/internal/websocket"
	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
)

var logs []LoggieLog

var websocketServer *wsHelper.WebsocketServer

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type LoggieLog struct {
	ID        int64     `json:"id"`
	Source    int64     `json:"source"`
	Severity  string    `json:"severity"`
	Category  string    `json:"category"`
	Resource  string    `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

func Run(serverPort int64, frontendPort int64) {
	fmt.Println("Starting server on port:", serverPort)

	websocketServer = wsHelper.NewServer()

	setupRoutes()
	handler := setupCors(frontendPort)

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", serverPort), *handler))
}

func setupRoutes() {
	http.HandleFunc("/", logEndpoint)
	http.HandleFunc("/ws", wsEndpoint)
}

func setupCors(frontendPort int64) *http.Handler {
	frontendAddr := fmt.Sprint("http://localhost:", frontendPort)

	origins := handlers.AllowedOrigins([]string{frontendAddr})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	headers := handlers.AllowedHeaders([]string{"Content-Type"})

	handler := handlers.CORS(origins, methods, headers)(http.DefaultServeMux)

	return &handler
}
