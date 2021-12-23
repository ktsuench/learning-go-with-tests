package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got, _ := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestConfigurableRacer(t *testing.T) {
	serverA := makeDelayedServer(1 * time.Second)
	serverB := makeDelayedServer(1 * time.Second)

	defer serverA.Close()
	defer serverB.Close()

	_, err := Racer(serverA.URL, serverB.URL, 1*time.Millisecond)

	if err == nil {
		t.Error("expected an error but didn't get one")
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		rw.WriteHeader(http.StatusOK)
	}))
}
