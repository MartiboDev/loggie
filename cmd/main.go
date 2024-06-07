package main

import (
	config "github.com/MartiboDev/loggie/internal/config"
	router "github.com/MartiboDev/loggie/internal/router"
)

func main() {
	config.Init()
	router.Run()
}
