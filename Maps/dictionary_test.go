package maps

import "testing"

func TestSearch(t *testing.T) {
	assertStrings := func (t testing.TB, got, want string)  {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	dictionary := Dictionary{"test":"This is test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is test"

		assertStrings(t, got, want)
	})

	// We expect to get err != nil
	t.Run("unknown word", func(t *testing.T) {
		_, gotErr := dictionary.Search("huihui")
		want := "word not found"

		if gotErr == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, gotErr.Error(), want)
	})
}