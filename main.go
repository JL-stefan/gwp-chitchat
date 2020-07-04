package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", service.index)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
}
