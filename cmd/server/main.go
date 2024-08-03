package server

import (
	"net/http"

	"github.com/XelerT/server.git/cmd/handelers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/update/`, handelers.ParsePost)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
