package clockface

import (
	"math"
	"testing"
	"time"
)

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 15), 0.5 * math.Pi},
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 45), 1.5 * math.Pi},
		{simpleTime(0, 0, 24), ((24.0 / 60.0) * 360) * (math.Pi / 180)},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := secondsInRadians(test.time)

			if got != test.angle {
				t.Fatalf("got %v, want %v rad", got, test.angle)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	angleAt48Sec := 48.0 * (math.Pi / 30)
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
		{simpleTime(0, 0, 48), Point{math.Sin(angleAt48Sec), math.Cos(angleAt48Sec)}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			if got != c.point {
				t.Fatalf("got %v, want %v", got, c.point)
			}
		})
	}
}
