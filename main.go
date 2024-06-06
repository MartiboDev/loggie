package main

import (
	api "github.com/MartiboDev/loggie/internal/api"
)

const serverPort int64 = 8080
const frontendPort int64 = 5173

func main() {
	api.Run(serverPort, frontendPort)
}
