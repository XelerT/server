package main

import (
	"fmt"
	"net/http"

	"github.com/XelerT/server.git/cmd/server/stor"
)

var storage *stor.MemStorage = stor.NewMemStorage()

func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	body += "Query parameters ===============\r\n"
	for k, v := range req.URL.Query() {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	res.Write([]byte(body))
}

func ParsePost(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(storage.Update(req.URL.String()))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/update/`, ParsePost)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
