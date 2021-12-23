package svg

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

func TestWriteSecondHand(t *testing.T) {
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
			Write(&b, test.time)

			svg := SVG{}
			// parse bytes buffer `b` as XML into SVG `svg`
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(test.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v in the SVG lines %+v", test.line, svg.Line)
			}
		})
	}
}

func TestWriteMinuteHand(t *testing.T) {
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
			Write(&b, test.time)

			svg := SVG{}
			// parse bytes buffer `b` as XML into SVG `svg`
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(test.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v in the SVG lines %+v", test.line, svg.Line)
			}
		})
	}
}

func TestWriteHourHand(t *testing.T) {
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
			Write(&b, test.time)

			svg := SVG{}
			// parse bytes buffer `b` as XML into SVG `svg`
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(test.line, svg.Line) {
				t.Errorf("Expected to find the hour hand line %+v in the SVG lines %+v", test.line, svg.Line)
			}
		})
	}
}

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

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
