package main

import "testing"

func TestHello( t *testing.T){
	assertCorrectMessage := func(t *testing.T, got, want string){
		t.Helper()
		if got != want{
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Saying hello to people", func(t *testing.T){
		got := Hello("Mark", "")
		want := "Hello, Mark"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World when an empty string is passed", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T){
		got := Hello("James", "Spanish")
		want := "Hola, James"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T){
		got := Hello("Claude", "French")
		want := "Bonjour, Claude"
		assertCorrectMessage(t, got, want)
	})
}