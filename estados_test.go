package main

import (
	"testing"
)

func TestEstados(t *testing.T) {
	m := Message{"", 0, "", 0}

	if buildMessage("/") != m {
		t.Error("Valor esperado {'', 0, '', 0}")
	}
}
