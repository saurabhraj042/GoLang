package sample1

import(
	"testing"
)

func TestRacer(t *testing.T) {
	slowUrl := "http://www.facebook.com"
	fastUrl := "http://www.quii.co.uk"

	got := Racer(slowUrl, fastUrl)
	want := fastUrl

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}