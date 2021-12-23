package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

// 	want := Point{X: 150, Y: 150 + 90}
// 	got := SecondHand(tm)

// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}
// }

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
		{simpleTime(0, 0, 15), 1.5 * math.Pi},
		{simpleTime(0, 0, 24), ((24 / 60) * 360) * (math.Pi / 180)},
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
