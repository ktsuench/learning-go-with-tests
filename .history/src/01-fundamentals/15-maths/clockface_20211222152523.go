package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

const (
	clockCenterX     = 150
	clockCenterY     = 150
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
)

const (
	svgStart = `<?xml version="1.0" encodin="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1/EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg" width="100%" height="100%" viewBox="0 0 300 300" version="2.0">`
	bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd = `</svg>`
)

func secondsInRadians(t time.Time) float64 {
	return float64(t.Second()) * math.Pi / 30
}

func minutesInRadians(t time.Time) float64 {
	minuteRatio := float64(t.Minute()) / 60
	return (minuteRatio*360)*(math.Pi/180) + secondsInRadians(t)/60
}

func hoursInRadians(t time.Time) float64 {
	hourRatio := float64(t.Hour()) / 24
	return (hourRatio*360)*(math.Pi/180) + minutesInRadians(t)/60 + secondsInRadians(t)/3600
}

func angleToPoint(angle float64) Point {
	x, y := math.Sin(angle), math.Cos(angle)
	return Point{x, y}
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return Point{}
}

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

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	io.WriteString(w, svgEnd)
}
