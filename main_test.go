package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(25, 35)

	if total != 60 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 60)
	}
}
