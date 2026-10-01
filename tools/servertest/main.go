package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	fmt.Println("starting...")

	client := &http.Client{Timeout: 5 * time.Second}

	tests := []struct {
		name   string
		method string
		path   string
	}{
		{
			name:   "post new order",
			method: "POST",
			path:   "/orders",
		},
		{
			name:   "get order id",
			method: "GET",
			path:   "/orders/1",
		},
		{
			name:   "get account",
			method: "GET",
			path:   "/accounts/2",
		},
		{
			name:   "publish nav for fund",
			method: "POST",
			path:   "/funds/1/navs",
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest(test.method, "http://localhost:8080"+test.path, nil)
		resp, _ := client.Do(req)

		body, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()

		fmt.Printf("%s: (%d), %s\n", test.name, resp.StatusCode, string(body))
	}
}
