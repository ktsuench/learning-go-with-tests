package clockface

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

type (
	Circle struct {
		Cx float64 `xml:"cx,attr"`
		Cy float64 `xml:"cy,attr"`
		R  float64 `xml:"r,attr"`
	}
	Line struct {
		X1 float64 `xml:"x1,attr"`
		Y1 float64 `xml:"y1,attr"`
		X2 float64 `xml:"x2,attr"`
		Y2 float64 `xml:"y2,attr"`
	}
	SVG struct {
		XMLName xml.Name `xml:"svg"`
		Xmlns   string   `xml:"xmlns,attr"`
		Width   string   `xml:"width,attr"`
		Height  string   `xml:"height,attr"`
		ViewBox string   `xml:"viewBox,attr"`
		Version string   `xml:"version,attr"`
		Circle  Circle   `xml:"circle"`
		Line    []Line   `xml:"line"`
	}
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

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
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

func TestSecondHandPoint(t *testing.T) {
	angleAt48Sec := 48 * (math.Pi / 30)
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 48), Point{math.Sin(angleAt48Sec), math.Cos(angleAt48Sec)}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("got %v, want %v", got, c.point)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	angleAt48Min := 48*(math.Pi/30) + 25*(math.Pi/1800)
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 15, 0), Point{1, 0}},
		{simpleTime(0, 48, 25), Point{math.Sin(angleAt48Min), math.Cos(angleAt48Min)}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minuteHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("got %v, want %v", got, c.point)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	angleAt8Hour := 8.2325 / 6 * math.Pi
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(15, 0, 0), Point{1, 0}},
		{simpleTime(8, 13, 57), Point{math.Sin(angleAt8Hour), math.Cos(angleAt8Hour)}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hourHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("got %v, want %v", got, c.point)
			}
		})
	}
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, test.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(test.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v in the SVG lines %+v", test.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 70}},
		{simpleTime(0, 30, 0), Line{150, 150, 150, 230}},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, test.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(test.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v in the SVG lines %+v", test.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 100}},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, test.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(test.line, svg.Line) {
				t.Errorf("Expected to find the hour hand line %+v in the SVG lines %+v", test.line, svg.Line)
			}
		})
	}
}
