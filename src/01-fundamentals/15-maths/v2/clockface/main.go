package main

import (
	"os"
	"time"

	"github.com/ktsuench/learning-go-with-tests/src/01-fundamentals/15-maths/v2/svg"
)

func main() {
	t := time.Now()

	svg.Write(os.Stdout, t)
}
