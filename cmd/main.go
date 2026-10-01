package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/laruibasar/nearfunds/internal/database"
	"github.com/laruibasar/nearfunds/internal/handler"
	"github.com/laruibasar/nearfunds/internal/processor"
)

func main() {
	fmt.Println("starting...")

	// Database setup.
	dbConfig, err := database.NewConfig()
	if err != nil {
		log.Fatalf("failed to config database: %v", err)
	}
	db, err := database.New(dbConfig)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	// Setup processor.
	proc := processor.New(db)

	handle := handler.New(proc)

	fmt.Println("...started")

	// TODO: fix the server setup for later, just to put it up.
	if err := http.ListenAndServe(":8080", handle.APIRoutes()); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
