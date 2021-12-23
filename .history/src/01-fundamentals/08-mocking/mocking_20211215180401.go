package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type DefaultSleeper struct{}

const finalWord = "Go!"
const countdownStart = 3

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(out, i)
	}

	s.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
