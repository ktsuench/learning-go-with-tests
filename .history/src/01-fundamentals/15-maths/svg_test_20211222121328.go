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
			want := test.angle

			if got != want {
				t.Fatalf("got %v, wanted %v rad", got, want)
			}
		})
	}
}
