package main

import (
	config "github.com/MartiboDev/loggie/config"
	router "github.com/MartiboDev/loggie/internal/router"
)

func main() {
	config.Init()
	router.Run()
}
