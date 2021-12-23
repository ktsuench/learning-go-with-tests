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

const center = 150
const secondHandLength = 90

func secondsInRadians(t time.Time) float64 {
	secondRatio := float64(t.Second()) / 60
	return (secondRatio * 360) * (math.Pi / 180)
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x, y := math.Sin(angle), math.Cos(angle)
	return Point{x, y}
}

// SecondHand is the unit vector of the second hand of
// an analogue clock at time `t` represented as a Point.
func SecondHand(t time.Time) Point {
	secondHand := secondHandPoint(t)
	x := center - math.Floor(secondHand.X)*secondHandLength
	y := center - math.Floor(secondHand.Y)*secondHandLength
	return Point{X: 150, Y: y}
}
