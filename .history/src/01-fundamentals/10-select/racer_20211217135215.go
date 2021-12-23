package main

import (
	"net/http"
)

func Racer(urlA, urlB string) (winner string) {
	// waits on both, but only runs case of one
	// that returns first
	select {
	case <-ping(urlA):
		return urlA
	case <-ping(urlB):
		return urlB
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
