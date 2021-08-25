package integers

import "testing"

// Function to test the Adder() function
func TestAdder(t *testing.T){
	got := Adder(2,2)
	want := 4

	if got != want{
		t.Errorf("got %d want %d", got, want)
	}
}