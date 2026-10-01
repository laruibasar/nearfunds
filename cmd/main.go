package main

import (
	"fmt"
	"log"
	"net/http"

	"git.sr.ht/~laruibasar/nearfunds/internal/handler"
	"git.sr.ht/~laruibasar/nearfunds/internal/processor"
)

func main() {
	fmt.Println("starting...")

	// Setup processor.
	proc := processor.New()

	handle := handler.New(proc)

	fmt.Println("...started")

	// TODO: fix the server setup for later, just to put it up.
	if err := http.ListenAndServe(":8080", handle.ApiRoutes()); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
