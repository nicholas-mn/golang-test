package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to the people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assetCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assetCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assetCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Nicholas", "French")
		want := "Bonjour, Nicholas"
		assetCorrectMessage(t, got, want)
	})
	t.Run("in Danish", func(t *testing.T) {
		got := Hello("Anna", "Danish")
		want := "Hej, Anna"
		assetCorrectMessage(t, got, want)
	})
}

func assetCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
