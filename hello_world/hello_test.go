package main

import "testing"

func TestHelloWorld(t *testing.T) {
	resultado := Hello("Mundo!")
	esperado := "Olá, Mundo!"

	if resultado != esperado {
		t.Errorf("Resultado '%s', esperado '%s'", resultado, esperado)
	}
}
