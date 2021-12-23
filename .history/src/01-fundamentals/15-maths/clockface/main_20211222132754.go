package main

import (
	"fmt"

	clockface "github.com/ktsuench/learning-go-with-tests/src/01-fundamentals/15-maths"
)

func secondHandTag(p clockface.Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func main() {

}
