package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

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
		{simpleTime(6, 0, 0), Line{150, 150, 150, 200}},
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
