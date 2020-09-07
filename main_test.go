package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := 2 + 2
	if resultado != 4 {
		t.Errorf("Soma esperada: %d, obtida: %d", 4, resultado)
	}
}
