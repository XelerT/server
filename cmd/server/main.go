package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{gauge: make(map[string]float64), counter: make(map[string]int64)}
}

var storage *MemStorage = NewMemStorage()

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

func (storage_ *MemStorage) updateCounter(name string, val int64) {
	_, ok := storage_.counter[name]
	if ok {
		storage_.counter[name] += val
	} else {
		storage_.counter[name] = val
	}
}

func (storage_ *MemStorage) updateGauge(name string, val float64) {
	storage_.gauge[name] = val
}

func (storage_ *MemStorage) Update(url string) int {
	parser := func(c rune) bool {
		return c == '/'
	}
	params := strings.FieldsFunc(url, parser)
	if len(params) != 4 {
		fmt.Println("not enoght params")
		fmt.Println(params)

		return http.StatusNotFound
	}
	metrType, name, val := params[1], params[2], params[3]

	if metrType == "counter" {
		convVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			fmt.Println("int can not convert")
			return http.StatusBadRequest
		}
		storage_.updateCounter(name, convVal)

		return http.StatusOK
	} else if metrType == "gauge" {
		convVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Println("float can not convert")

			return http.StatusBadRequest
		}
		storage_.updateGauge(name, convVal)

		return http.StatusOK
	}

	fmt.Println("wrong type")
	return http.StatusBadRequest
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
