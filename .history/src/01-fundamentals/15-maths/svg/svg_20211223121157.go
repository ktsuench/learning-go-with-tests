package svg

import (
	"fmt"
	"io"
	"time"
)

func makeHand(p Point, length float64) Point {
	p.X = clockCenterX - p.X*length
	p.Y = clockCenterX - p.Y*length
	return p
}

// secondHand is the unit vector of the second hand of
// an analogue clock at time `t` represented as a Point.
func secondHand(w io.Writer, t time.Time) {
	p := makeHand(secondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.2f" y2="%.2f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(minuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.2f" y2="%.2f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(hourHandPoint(t), hourHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.2f" y2="%.2f" style="fill:none;stroke:#000;stroke-width:5px;"/>`, p.X, p.Y)
}

func Write(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}
