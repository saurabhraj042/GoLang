package arrays

import (
	"reflect"
	"testing"
)

// Test for Sum() function of arrays
func TestSum(t *testing.T){
	assertCorrectMessage := func(t testing.TB, got, want int){
		t.Helper()
		
		if got != want {
			t.Errorf("got %d want %d ", got, want)
		}
	}

	// Subtest for array of fixed size
	t.Run("collection of fixed size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, got, want)
	})

	// Subtest for variable size array
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want)
	})
}

// Test for SumAll() function
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{2, 3})
	want := []int{3, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Test for SumTails() function
func TestSumTails(t *testing.T){
	got := SumTails([]int{1, 2, 3, 4}, []int{2, 3})
	want := []int{9, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}