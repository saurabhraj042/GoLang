package maps

import "testing"

func TestSearch(t *testing.T) {
	assertStrings := func (t testing.TB, got, want string)  {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test":"This is test"}
		got := dictionary.Search("test")
		want := "This is test"

		assertStrings(t, got, want)
	})
}