package main

import (
	"net/http"
	"fmt"
)

func main() {
	server := http.Server{Addr: ":8080", Handler: http.NewServeMux()}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Errorf("Cannot connect to server %v", err)
	}
}