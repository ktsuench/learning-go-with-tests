package main

import (
	"net/http"
)

func Racer(urlA, urlB string) (url string) {
	aDuration := measureResponseTime(urlA)
	bDuration := measureResponseTime(urlB)

	if aDuration < bDuration {
		url = urlA
	} else {
		url = urlB
	}

	return
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
