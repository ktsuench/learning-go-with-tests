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

// SecondHand is the unit vector of the second hand of
// an analogue clock at time `t` represented as a Point.
func SecondHand(t time.Time) Point {
	y := 150 - math.Floor(math.Cos((float64(t.Second())/60)*(math.Pi/180))*60)
	return Point{X: 150, Y: y}
}
