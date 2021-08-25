package iteration

import (
	"fmt"
	"testing"
)

// Test for Repeat() function
func TestRepeat(t *testing.T) {
	got := Repeat("a",10)
	want := "aaaaaaaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Benchmark for Repeat() function
func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// Example for Repeat() function
func ExampleRepeat() {
	got := Repeat("a",2)
	fmt.Println(got)
	// Output: aa
}