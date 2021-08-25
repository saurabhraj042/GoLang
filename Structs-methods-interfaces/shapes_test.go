package structsmethodsinterfaces

import "testing"

// Test for Perimeter() function
func TestPerimeter(t *testing.T){
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Test for Area() function
func TestArea(t *testing.T){
	rectangle := Rectangle{10.0, 10.0}
	got := Area(rectangle)
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}