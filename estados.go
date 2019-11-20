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
	case "/AM":
		m = Message{"Amazonas", 1570745.7, "Manaus", 4063614}
	case "/BA":
		m = Message{"Bahia", 564692.7, "Salvador", 15344447}
	case "/CE":
		m = Message{"Ceará", 148825.6, "Fortaleza", 4063614}
	case "/DF":
		m = Message{"Distrito Federal", 5822.1, "Brasília", 3039444}
	case "/ES":
		m = Message{"Espirito Santo", 46077.5, "Vitória", 4016356}
	case "/GO":
		m = Message{"Goiás", 340086.7, "Goiânia", 6778772}
	case "/MA":
		m = Message{"Maranhão", 331983.3, "São Luís", 7000229}
	case "/MT":
		m = Message{"Mato Grosso", 903357.9, "Cuiabá", 3344544}
	case "/MS":
		m = Message{"Mato Grosso do Sul", 357125.0, "Campo Grande", 2713147}
	case "/MG":
		m = Message{"Minas Gerais", 58652837, "Belo Horizonte", 21119536}
	case "/PA":
		m = Message{"Pará", 1247689.5, "Belém", 8366628}
	case "/PB":
		m = Message{"Paraíba", 56439.8, "João Pessoa", 4025558}
	case "/PR":
		m = Message{"Paraná", 199314.9, "Curituba", 11320892}
	case "/PE":
		m = Message{"Pernambuco", 98311.6, "Recife", 9473266}
	case "/PI":
		m = Message{"Piauí", 251529.2, "Teresina", 3219257}
	case "/RJ":
		m = Message{"Rio de Janeiro", 43696.1, "Rio de Janeiro", 16718956}
	case "/RN":
		m = Message{"Rio Grande do Norte", 52796.8, "Natal", 3507003}
	case "/RS":
		m = Message{"Rio Grande do Sul", 281748.5, "Porto Alegre", 11322895}
	case "/RO":
		m = Message{"Rondônia", 237576.2, "Porto Velho", 1805788}
	case "/RR":
		m = Message{"Roraima", 224299.0, "Boa Vista", 522636}
	case "/SC":
		m = Message{"Santa Catarina", 95346.2, "Florianópolis", 7001161}
	case "/SP":
		m = Message{"São Paulo", 43696.1, "São Paulo", 16718956}
	case "/SE":
		m = Message{"Sergipe", 1570745.7, "Aracaju", 4063614}
	case "/TO":
		m = Message{"Tocantins", 1570745.7, "Palmas", 4063614}
	default:
		w.WriteHeader(400) // Return 400 Bad Request.
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}

}

func main() {

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
