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
	return math.Pi
}

// SecondHand is the unit vector of the second hand of
// an analogue clock at time `t` represented as a Point.
func SecondHand(t time.Time) Point {
	y := 150 - math.Floor(math.Cos((float64(t.Second())/60)*360*(math.Pi/180))*90)
	return Point{X: 150, Y: y}
}
