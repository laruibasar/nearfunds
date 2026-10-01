package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/laruibasar/nearfunds/internal/handler"
	"github.com/laruibasar/nearfunds/internal/processor"
)

func main() {
	fmt.Println("starting...")

	// Setup processor.
	proc := processor.New()

	handle := handler.New(proc)

	fmt.Println("...started")

	// TODO: fix the server setup for later, just to put it up.
	if err := http.ListenAndServe(":8080", handle.APIRoutes()); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
