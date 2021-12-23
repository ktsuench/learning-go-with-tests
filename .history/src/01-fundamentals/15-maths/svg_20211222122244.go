package clockface

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	secondRatio := float64(t.Second()) / 60
	return (secondRatio * 360) * (math.Pi / 180)
}

func secondHandPoint(t time.Time) Point {
	return Point{0, -1}
}

// SecondHand is the unit vector of the second hand of
// an analogue clock at time `t` represented as a Point.
func SecondHand(t time.Time) Point {
	y := 150 - math.Floor(math.Cos(secondsInRadians(t))*90)
	return Point{X: 150, Y: y}
}
