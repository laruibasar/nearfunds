package main

import (
	"fmt"
	"log"
	"net/http"

	"git.sr.ht/~laruibasar/nearfunds/internal/handler"
)

func main() {
	fmt.Println("starting...")

	handle := handler.New()

	fmt.Println("...started")

	// TODO: fix the server setup for later, just to put it up.
	if err := http.ListenAndServe(":8080", handle.ApiRoutes()); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
