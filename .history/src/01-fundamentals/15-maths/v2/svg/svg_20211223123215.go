package svg

import (
	"fmt"
	"io"
	"time"

	cf "github.com/ktsuench/learning-go-with-tests/src/01-fundamentals/15-maths/v2"
)

// Outputs an analog clock showing time `t` as a SVG to writer `w`
func Write(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

const (
	clockCenterX     = 150
	clockCenterY     = 150
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
)

// secondHand is the unit vector of the second hand of
// an analogue clock at time `t` represented as a Point.
func secondHand(w io.Writer, t time.Time) {
	p := makeHand(cf.SecondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.2f" y2="%.2f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(cf.MinuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.2f" y2="%.2f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(cf.HourHandPoint(t), hourHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.2f" y2="%.2f" style="fill:none;stroke:#000;stroke-width:5px;"/>`, p.X, p.Y)
}

func makeHand(p cf.Point, length float64) cf.Point {
	p.X = clockCenterX - p.X*length
	p.Y = clockCenterX - p.Y*length
	return p
}

const (
	svgStart = `<?xml version="1.0" encodin="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1/EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg" width="100%" height="100%" viewBox="0 0 300 300" version="2.0">`
	bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd = `</svg>`
)
