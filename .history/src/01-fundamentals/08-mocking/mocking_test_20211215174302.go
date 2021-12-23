package main

import (
	"bytes"
	"reflect"
	"testing"
)

const (
	write = "write"
	sleep = " sleep"
)

type SpyCountdownOps struct {
	Calls []string
}

func (s *SpyCountdownOps) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOps) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCoundown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := bytes.Buffer{}
		SpySleeper := &SpyCountdownOps{}

		Countdown(&buffer, SpySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if calls := len(SpySleeper.Calls); calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", calls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOps{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{}
		for i := 0; i < 4; i++ {
			want = append(want, sleep, write)
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
