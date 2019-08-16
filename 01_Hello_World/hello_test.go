package main

import "testing"

func TestHello(t *testing.T) {
	assert := func(t *testing.T, received, expected string) {
		// t.Helper() is needed to tell the test suite that this method is a helper
		// By doing this when it fails the line number reported will be in our
		// function call rather than inside our test helper
		t.Helper()
		if received != expected {
			t.Errorf("❌ received %q expected %q", received, expected)
		}
	}

	t.Run("should say hello to people", func(t *testing.T) {
		received := Hello("Andre", "")
		expected := "Hello, Andre!"

		assert(t, received, expected)
	})

	t.Run("should say 'Hello, Stranger!' when an empty string is provided", func(t *testing.T) {
		received := Hello("", "")
		expected := "Hello, Stranger!"

		assert(t, received, expected)
	})

	t.Run("should say 'Olá, Estranho!' when called with 'PT' as second parameter", func(t *testing.T) {
		received := Hello("", "PT")
		expected := "Olá, Estranho!"

		assert(t, received, expected)
	})

	t.Run("should say 'Bonjour, Étranger!' when called with 'FR' as second parameter", func(t *testing.T) {
		received := Hello("", "FR")
		expected := "Bonjour, Étranger!"

		assert(t, received, expected)
	})
}
