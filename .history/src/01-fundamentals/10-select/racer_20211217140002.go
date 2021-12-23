package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(urlA, urlB string) (winner string, err error) {
	return ConfigurableRacer(urlA, urlB, tenSecondTimeout)
}

func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (winner string, err error) {
	// waits on both, but only runs case of one
	// that returns first
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %q and %q", urlA, urlB)
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
