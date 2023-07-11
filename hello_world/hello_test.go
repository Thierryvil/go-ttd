package main

import "testing"

func TestHelloWorld(t *testing.T) {

	checkRightAnswer := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result: '%s', expected: '%s'", result, expected)

		}
	}
	t.Run("Should be return: Olá, Chris", func(t *testing.T) {
		result := Hello("Chris!", "")
		expected := "Olá, Chris!"
		checkRightAnswer(t, result, expected)

		t.Run("Should be return: Olá, Mundo!", func(t *testing.T) {
			result := Hello("Mundo!", "")
			expected := "Olá, Mundo!"

			checkRightAnswer(t, result, expected)
		})
	})

	t.Run("Should be return : Holá, Eddie", func(t *testing.T) {
		result := Hello("Eddie!", "spanish")
		expected := "Hola, Eddie!"
		checkRightAnswer(t, result, expected)
	})

	t.Run("Should be return Bonjour, Tobias!", func(t *testing.T) {
		result := Hello("Tobias!", "french")
		expected := "Bonjour, Tobias!"
		checkRightAnswer(t, result, expected)
	})
}
