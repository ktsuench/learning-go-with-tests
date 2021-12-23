package clockface

import (
	"math"
	"testing"
	"time"
)

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 15), 0.5 * math.Pi},
		{simpleTime(0, 0, 24), 0.8 * math.Pi},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := secondsInRadians(test.time)
			if !roughlyEqualFloat64(got, test.angle) {
				t.Fatalf("got %v, want %v rad", got, test.angle)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 15, 0), 0.5 * math.Pi},
		{simpleTime(0, 24, 18), 0.81 * math.Pi},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := minutesInRadians(test.time)
			if !roughlyEqualFloat64(got, test.angle) {
				t.Fatalf("got %v, want %v rad", got, test.angle)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(6, 54, 27), 1.15125 * math.Pi},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := hoursInRadians(test.time)
			if !roughlyEqualFloat64(got, test.angle) {
				t.Fatalf("got %v, want %v rad", got, test.angle)
			}
		})
	}
}
