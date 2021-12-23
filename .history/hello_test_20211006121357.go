package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMsg := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMsg(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMsg(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMsg(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Sylvie", "French")
		want := "Bonjour, Sylvie"
		assertCorrectMsg(t, got, want)
	})
}
