package clockface

import (
	"math"
	"time"
)

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
