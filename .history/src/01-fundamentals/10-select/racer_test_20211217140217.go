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

	if err != nil {
		t.Fatalf("did not expect an error but got one %v", err)
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestConfigurableRacer(t *testing.T) {
	server := makeDelayedServer(1 * time.Second)

	defer server.Close()

	_, err := ConfigurableRacer(server.URL, server.URL, 1*time.Millisecond)

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
