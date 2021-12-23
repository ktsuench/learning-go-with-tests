package main

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (url string) {
	startA := time.Now()
	http.Get(urlA)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(urlB)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		url = urlA
	} else {
		url = urlB
	}

	return
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
