package main

import "testing"

func TestHello(t *testing.T){
	got := Hello()
	want := "Hello World from outside the main() scope"

	if got != want{
		t.Errorf("got %q want %q",got,want)
	}
}