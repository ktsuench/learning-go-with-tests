package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(out, "%d", i)
	}
	fmt.Fprint(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}