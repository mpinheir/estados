package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, req.URL.Path)

	m := Message{"Alice", "Hello", 1294706395881547000}

	b, err := json.Marshal(m)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, b)

}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
