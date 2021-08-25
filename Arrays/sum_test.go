package arrays

import "testing"

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
		numbers := [5]int{1, 2, 3, 4, 5}

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