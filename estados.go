package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

type dadosUf struct {
	Estado, Capital string
	Area            float64
	Populacao       int64
}

func puxarUFs() map[string]dadosUf {
	ufs := map[string]dadosUf{}

	c := colly.NewCollector()
	c.OnHTML("#mwAfg table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, line *colly.HTMLElement) {
			var uf dadosUf
			var abreviacao string
			line.ForEach("td", func(_ int, h *colly.HTMLElement) {
				i := h.Index
				if i == 1 {
					uf.Estado = h.Text
				}
				if i == 2 {
					abreviacao = h.Text
				}
				if i == 3 {
					uf.Capital = h.Text
				}
				if i == 4 {
					sanitizado := strings.ReplaceAll(h.Text, "\u00a0", "")
					comPonto := strings.ReplaceAll(sanitizado, ",", ".")
					f, err := strconv.ParseFloat(comPonto, 64)
					if err != nil {
						log.Fatal(err)
					}
					uf.Area = f
				}
				if i == 5 {
					i, err := strconv.ParseInt(strings.ReplaceAll(h.Text, "\u00a0", ""), 10, 64)
					if err != nil {
						log.Fatal(err)
					}
					uf.Populacao = i
					ufs[strings.ToUpper(abreviacao)] = uf
				}
			})
		})
	})

	err := c.Visit("https://pt.wikipedia.org/wiki/Unidades_federativas_do_Brasil")
	if err != nil {
		log.Fatal(err)
	}

	return ufs
}

func dadosUfHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jsonResponse []byte
	vars := mux.Vars(r)
	uf := vars["UF"]

	allUfs := puxarUFs()
	if strings.EqualFold(uf, "todos") {
		jsonResponse, _ = json.Marshal(allUfs)
	} else {
		jsonResponse, _ = json.Marshal(allUfs[strings.ToUpper(uf)])
	}
	_, _ = w.Write(jsonResponse)
}
