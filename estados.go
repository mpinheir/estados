package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Estado    string
	Area      float64
	Capital   string
	Populacao int64
}

func handler(w http.ResponseWriter, req *http.Request) {
	m := Message{"", 0, "", 0}

	switch strings.ToUpper(req.URL.Path) {
	case "/AC":
		m = Message{"Acre", 152581.4, "Rio Branco", 829619}
	case "/AL":
		m = Message{"Alagoas", 27767.7, "Maceió", 3375823}
	case "/AP":
		m = Message{"Amapá", 142814.6, "Macapá", 797722}
	case "/RJ":
		m = Message{"Rio de Janeiro", 43696.1, "Rio de Janeiro", 16718956}
	case "/SP":
		m = Message{"São Paulo", 43696.1, "São Paulo", 16718956}
	default:
		//needs to return message regarding invalid format
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)

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
