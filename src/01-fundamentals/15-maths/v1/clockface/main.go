package main

import (
	"os"
	"time"

	clockface "github.com/ktsuench/learning-go-with-tests/src/01-fundamentals/15-maths/v1"
)

func main() {
	t := time.Now()

	clockface.SVGWriter(os.Stdout, t)
}
