package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	clockface "github.com/ktsuench/learning-go-with-tests/src/01-fundamentals/15-maths"
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

// func TestSecondHandAtMidnight(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

// 	want := clockface.Point{X: 150, Y: 150 - 90}
// 	got := clockface.SecondHand(tm)

// 	if got != want {
// 		t.Errorf("got %v, want %v", got, want)
// 	}
// }

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

// 	want := clockface.Point{X: 150, Y: 150 + 90}
// 	got := clockface.SecondHand(tm)

// 	if got != want {
// 		t.Errorf("got %v, want %v", got, want)
// 	}
//}

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	b := bytes.Buffer{}
	clockface.SVGWriter(&b, tm)

	svg := SVG{}
	xml.Unmarshal(b.Bytes(), &svg)

	x2, y2 := "150.000", "60.000"

	for _, line := range svg.Line {
		if line.X2 == x2 && line.Y2 == y2 {
			return
		}
	}

	t.Errorf("Expected to find the second hand with x2 of %+v and y2 of %+v in the SVG output %v", x2, y2, b.String())
}
