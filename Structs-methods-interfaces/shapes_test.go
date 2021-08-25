package structsmethodsinterfaces

import "testing"

// Test for Perimeter() function
func TestPerimeter(t *testing.T){
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}