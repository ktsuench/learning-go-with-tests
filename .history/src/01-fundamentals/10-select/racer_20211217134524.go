package main

import (
	"net/http"
	"time"
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

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
