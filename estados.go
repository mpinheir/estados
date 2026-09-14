package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type dadosUf struct {
	Estado, Capital string
	Area            float64
	Populacao       int64
}

var allUfs = map[string]dadosUf{
	"AC": {"Acre", "Rio Branco", 152581.4, 829619},
	"AL": {"Alagoas", "Maceió", 27767.7, 3375823},
	"AP": {"Amapá", "Macapá", 142814.6, 797722},
	"AM": {"Amazonas", "Manaus", 1570745.7, 4063614},
	"BA": {"Bahia", "Salvador", 564692.7, 15344447},
	"CE": {"Ceará", "Fortaleza", 148825.6, 4063614},
	"DF": {"Distrito Federal", "Brasília", 5822.1, 3039444},
	"ES": {"Espirito Santo", "Vitória", 46077.5, 4016356},
	"GO": {"Goiás", "Goiânia", 340086.7, 6778772},
	"MA": {"Maranhão", "São Luís", 331983.3, 7000229},
	"MT": {"Mato Grosso", "Cuiabá", 903357.9, 3344544},
	"MS": {"Mato Grosso do Sul", "Campo Grande", 357125.0, 2713147},
	"MG": {"Minas Gerais", "Belo Horizonte", 58652837., 21119536},
	"PA": {"Pará", "Belém", 1247689.5, 8366628},
	"PB": {"Paraíba", "João Pessoa", 56439.8, 4025558},
	"PR": {"Paraná", "Curitiba", 199314.9, 11320892},
	"PE": {"Pernambuco", "Recife", 98311.6, 9473266},
	"PI": {"Piauí", "Teresina", 251529.2, 3219257},
	"RJ": {"Rio de Janeiro", "Rio de Janeiro", 43696.1, 16718956},
	"RN": {"Rio Grande do Norte", "Natal", 52796.8, 3507003},
	"RS": {"Rio Grande do Sul", "Porto Alegre", 281748.5, 11322895},
	"RO": {"Rondônia", "Porto Velho", 237576.2, 1805788},
	"RR": {"Roraima", "Boa Vista", 224299.0, 522636},
	"SC": {"Santa Catarina", "Florianópolis", 95346.2, 7001161},
	"SP": {"São Paulo", "São Paulo", 43696.1, 16718956},
	"SE": {"Sergipe", "Aracaju", 1570745.7, 4063614},
	"TO": {"Tocantins", "Palmas", 1570745.7, 4063614},
}

func dadosUfHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jsonResponse []byte
	vars := mux.Vars(r)
	uf := vars["UF"]
	if strings.EqualFold(uf, "todos") {
		jsonResponse, _ = json.Marshal(allUfs)
	} else {
		jsonResponse, _ = json.Marshal(allUfs[strings.ToUpper(uf)])
	}
	_, _ = w.Write(jsonResponse)
}
