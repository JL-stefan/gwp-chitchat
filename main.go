package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JL-stefan/gwp-chitchat/data"
	"github.com/prometheus/common/log"
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

type Thread struct {
	ID        int    `json:"id"`
	Uuid      string `json:"uuid"`
	Topic     string `json:"topic"`
	UserID    int    `json:"userID"`
	CreatedAt int64  `json:"createdAt"`
}

func index(w http.ResponseWriter, r *http.Request) {
	var err error
	res, err := data.Threads()
	var result = make([]Thread, len(res))
	// var result = make([]Thread, 0)
	log.Info(result)
	if err != nil {
		return
	}
	for i, value := range res {
		result[i].ID = value.ID
		result[i].Uuid = value.Uuid
		result[i].Topic = value.Topic
		result[i].UserID = value.UserID
		result[i].CreatedAt = value.CreatedAt.Unix()
	}
	log.Info(result)
	ret, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Info(string(ret))
	fmt.Fprintf(w, "Hello World\n%v", string(ret))
	// fmt.Fprint(w, "Hello World")
	return
}
