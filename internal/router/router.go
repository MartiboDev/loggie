package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MartiboDev/loggie/config"
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

func Run() {
	serverPort := config.Get("SERVER_PORT")

	fmt.Println("Starting server on port:", serverPort)

	websocketServer = wsHelper.NewServer()

	setupRoutes()
	handler := setupCors()

	http.ListenAndServe(fmt.Sprint(":", serverPort), *handler)
}

func setupRoutes() {
	http.HandleFunc("/", logEndpoint)
	http.HandleFunc("/ws", wsEndpoint)
}

func setupCors() *http.Handler {
	frontendPort := config.Get("FRONTEND_PORT")

	frontendAddr := fmt.Sprint("http://localhost:", frontendPort)

	origins := handlers.AllowedOrigins([]string{frontendAddr})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	headers := handlers.AllowedHeaders([]string{"Content-Type"})

	handler := handlers.CORS(origins, methods, headers)(http.DefaultServeMux)

	return &handler
}
