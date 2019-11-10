package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Estado string
	Capital string
	Populacao int64
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, req.URL.Path)

	m := Message{"Rio de Janeiro", "Rio de Janeiro", 8000000}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}

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
	//http.HandleFunc("/headers", headers)

	port := os.Getenv("PORT")
	if port == "" {
			port = "8080"
			log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
