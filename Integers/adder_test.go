package integers

import (
	"fmt"
	"testing"
)

// Function to test the Adder() function
func TestAdder(t *testing.T){
	got := Adder(2,2)
	want := 4

	if got != want{
		t.Errorf("got %d want %d", got, want)
	}
}

// Example for Adder() function
func ExampleAdder() {
	got := Adder(3, 2)
	fmt.Println(got)
	// Output: 5
}