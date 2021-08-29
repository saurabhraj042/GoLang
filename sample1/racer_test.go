package sample1

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		rw.WriteHeader(http.StatusOK)
	}))
	
	fastServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		
		rw.WriteHeader(http.StatusOK)
	}))
	
	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	got := Racer(slowUrl, fastUrl)
	want := fastUrl

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}