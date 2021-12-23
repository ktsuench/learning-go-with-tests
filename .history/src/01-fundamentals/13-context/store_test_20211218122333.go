package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	t         *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) string {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <- ctx.Done():
				s.t.Log("spy store got cancelled"
			default:
			time.Sleep(10*time.Millisecond))
			}
		}
	}
}

func TestServer(t *testing.T) {
	t.Run("run normally", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q, want %q", response.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{delay: 100 * time.Millisecond, response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.assertWasCancelled()
	})
}
