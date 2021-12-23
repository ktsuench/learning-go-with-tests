package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	delay     time.Duration
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(s.delay)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("run normally", func(t *testing.T) {
		data := "hello world"
		svr := Server(&SpyStore{response: data})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q, want %q", response.Body.String(), data)
		}
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {})
}
