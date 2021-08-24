package main

import "testing"

func TestHello(t *testing.T){
	got := Hello()
	want := "Hello World from outside the main() scope"

	if got != want{
		t.Errorf("got %q want %q",got,want)
	}
}

func TestHelloAgain(t *testing.T){
	assertCorrectMessage := func(t testing.TB, got, want string){
		t.Helper()
		
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloAgain("Raj")
		want := "Hello, Raj"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := HelloAgain("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}