// Package handler will have the http server setup and api routes.
package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/laruibasar/nearfunds/internal/models"
	"github.com/laruibasar/nearfunds/internal/processor"
)

const headerCorrelation = "Idempotency-Key"

// push dependencies here, like the database.
type server struct {
	processor processor.Processor
}

func New(p processor.Processor) *server {
	return &server{processor: p}
}

func (s *server) APIRoutes() http.Handler {
	srv := http.NewServeMux()

	srv.HandleFunc("POST /orders", s.handleOrders)
	srv.HandleFunc("GET /orders", s.handleSearchOrders)
	srv.HandleFunc("GET /orders/{id}", s.handleGetOrder)
	srv.HandleFunc("POST /orders/{id}/cancel", s.handleOrderCancel)
	srv.HandleFunc("GET /orders/{id}/events", s.handleGetOrderEvents)
	srv.HandleFunc("GET /accounts/{id}", s.handleGetAccount)
	srv.HandleFunc("POST /funds/{id}/navs", s.handlePublishNAV)
	srv.HandleFunc("GET /healthz", s.handleCheckHealth)
	srv.HandleFunc("GET /readiness", s.handleCheckReadiness)

	return srv
}

func (s *server) handleOrders(w http.ResponseWriter, r *http.Request) {
	correlationID := r.Header.Get(headerCorrelation)
	if correlationID != "" {
		// This is mandatory, throw an 412, the header is missing and needed.
		err := errors.New("missing order identifier")

		http.Error(w, err.Error(), http.StatusPreconditionFailed)

		return
	}

	// Extract body from request and convert into data model from app.
	var order models.Order
	jsonDecoded := json.NewDecoder(r.Body)
	if err := jsonDecoded.Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	correlationUUID := uuid.MustParse(correlationID)
	order.Correlation = correlationUUID

	err := s.processor.CreateOrder(order, time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Println("order created ...to implement")
}

func (s *server) handleSearchOrders(w http.ResponseWriter, r *http.Request) {
	hasSearch := len(r.URL.Query()) > 0

	if !hasSearch {
		fmt.Println("full list orders but limit to user in authentication... to implement")

		return
	}

	fmt.Printf("search order with parameters: to implement")

	w.WriteHeader(http.StatusOK)
}

func (s *server) handleGetOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("get order received: %s ...to implement", r.PathValue("id"))

	w.WriteHeader(http.StatusOK)
}

func (s *server) handleOrderCancel(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("cancel order: %s ...to implement", r.PathValue("id"))

	w.WriteHeader(http.StatusAccepted)
}

func (s *server) handleGetOrderEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("list events for order: %s ...to implement", r.PathValue("id"))

	w.WriteHeader(http.StatusOK)
}

func (s *server) handleGetAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("get account: %s ...to implement", r.PathValue("id"))

	w.WriteHeader(http.StatusOK)
}

func (s *server) handlePublishNAV(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("publish NAV value for fund: %s ...to implement", r.PathValue("id"))

	w.WriteHeader(http.StatusAccepted)
}

func (s *server) handleCheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *server) handleCheckReadiness(w http.ResponseWriter, r *http.Request) {
	// TODO: Check if other services are okay.
	w.WriteHeader(http.StatusOK)
}
