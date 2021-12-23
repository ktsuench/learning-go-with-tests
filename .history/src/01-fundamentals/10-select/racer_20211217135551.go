package main

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string, err error) {
	// waits on both, but only runs case of one
	// that returns first
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(10 * time.Second):
		return "", error.New()
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		// closing the channel will return a signal
		close(ch)
	}()
	return ch
}
