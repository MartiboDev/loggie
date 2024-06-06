package loggie

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port:", Port)

	SetupRoutes()
	handler := SetupCors()

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", Port), *handler))
}
