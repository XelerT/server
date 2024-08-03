package handelers

import (
	"fmt"
	"net/http"

	"github.com/XelerT/server.git/cmd/stor"
)

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
	res.WriteHeader(stor.Storage.Update(req.URL.String()))
}
