package main

import (
	"fmt"
	"net/http"

	"github.com/JL-stefan/gwp-chitchat/data"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	res := data.Threads()
	fmt.Fprint(w, "Hello World\n", data.Threads())
}
