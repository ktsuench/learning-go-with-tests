package main

import (
	"bytes"
	"testing"
)

func TestCoundown(t *testing.T) {
	buffer := bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.ErrorF("got %q want %q", got, want)
	}
}
