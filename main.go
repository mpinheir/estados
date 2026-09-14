package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{UF}", dadosUfHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Assumindo porta %s comp padrão", port)
	} else {
		log.Printf("Utilizando porta %s", port)
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
