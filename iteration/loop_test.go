package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a",10)
	want := "aaaaaaaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	got := Repeat("a",2)
	fmt.Println(got)
	// Output: aa
}