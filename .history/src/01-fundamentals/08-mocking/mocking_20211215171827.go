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

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	s.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
