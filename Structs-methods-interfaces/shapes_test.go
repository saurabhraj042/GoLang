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

// Test for Area() function
func TestArea(t *testing.T){
	got := Area(10.0, 10.0)
	want := 200.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}