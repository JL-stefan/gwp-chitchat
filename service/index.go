package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JL-stefan/gwp-chitchat/data"
	"github.com/prometheus/common/log"
)

func index(w http.ResponseWriter, r *http.Request) {
	var err error
	res, err := data.Threads()
	if err != nil {
		return
	}
	ret, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Info(string(ret))
	fmt.Fprintf(w, "Hello World\n%v", string(ret))
	// fmt.Fprint(w, "Hello World")
	return
}
