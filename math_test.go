package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("o resultado da da soma é invalido resultado %d esperado %d", total, 30)
	}
}