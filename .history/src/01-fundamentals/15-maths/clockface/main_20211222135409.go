package main

import (
	"os"
	"time"

	clockface "github.com/ktsuench/learning-go-with-tests/src/01-fundamentals/15-maths"
)

func main() {
	t := time.Now()

	clockface.SVGWriter(os.Stdout, t)
}
