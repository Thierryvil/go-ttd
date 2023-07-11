package main

import "testing"

func TestHelloWorld(t *testing.T) {

	checkRightAnswer := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("resultado: '%s', esperado: '%s'", result, expected)

		}
	}
	t.Run("Deve retornar: Olá, Chris", func(t *testing.T) {
		result := Hello("Chris!")
		expected := "Olá, Chris!"
		checkRightAnswer(t, result, expected)

		t.Run("Deve retornar Olá, Mundo!", func(t *testing.T) {
			result := Hello("Mundo!")
			expected := "Olá, Mundo!"

			checkRightAnswer(t, result, expected)
		})
	})
}
