// Package handler will have the http server setup and api routes.
package handler

import (
	"fmt"
	"net/http"
)

// push dependencies here, like the database.
type server struct{}

func New() *server {
	return &server{}
}

func (s *server) ApiRoutes() http.Handler {
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
