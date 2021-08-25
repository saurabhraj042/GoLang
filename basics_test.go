package main

import "testing"

// Test for Hello() function
func TestHello(t *testing.T){
	got := Hello()
	want := "Hello World from outside the main() scope"

	if got != want{
		t.Errorf("got %q want %q",got,want)
	}
}

// Test for HelloAgain() function
func TestHelloAgain(t *testing.T){
	assertCorrectMessage := func(t testing.TB, got, want string){
		t.Helper()
		
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	
	// Subtest for Default behaviour
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloAgain("Raj", "")
		want := "Hello, Raj"

		assertCorrectMessage(t, got, want)
	})

	// Subtest to check for Empty inputs
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := HelloAgain("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	// Subtest to check if Language parameter is passed
	t.Run("in French", func(t *testing.T) {
		got := HelloAgain("Raj", "French")
		want := "Bonjour, Raj"

		assertCorrectMessage(t, got, want)
	})
}