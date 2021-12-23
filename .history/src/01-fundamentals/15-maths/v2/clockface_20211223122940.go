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
	return (float64(t.Second()) / 30) * math.Pi
}

func minutesInRadians(t time.Time) float64 {
	minuteRad := (float64(t.Minute()) / 30) * math.Pi
	return minuteRad + secondsInRadians(t)/60
}

func hoursInRadians(t time.Time) float64 {
	hourRatio := (float64(t.Hour()) / 6) * math.Pi
	return hourRatio + minutesInRadians(t)/12
}

func angleToPoint(angle float64) Point {
	x, y := math.Sin(angle), math.Cos(angle)
	return Point{x, y}
}

func SecondHandPoint(t time.Time) Point {
	return angleToPoint(cf.secondsInRadians(t))
}

func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(cf.minutesInRadians(t))
}

func HourHandPoint(t time.Time) Point {
	return angleToPoint(cf.hoursInRadians(t))
}
