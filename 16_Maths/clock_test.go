package maths

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, test := range cases {
		t.Run("should be able to calculate the radians for "+formatDate(test.time), func(t *testing.T) {
			received := secondsInRadians(test.time)
			expected := test.angle

			if !approximatedFloat64(received, expected) {
				t.Fatalf("❌ expected %v received %v", expected, received)
			}
		})
	}

}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), (math.Pi / (30 * 60)) * 7},
	}

	for _, test := range cases {
		t.Run("should be able to calculate the radians for "+formatDate(test.time), func(t *testing.T) {
			received := minutesInRadians(test.time)
			expected := test.angle

			if !approximatedFloat64(received, expected) {
				t.Fatalf("❌ expected %v received %v", expected, received)
			}
		})
	}

}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, test := range cases {
		t.Run("should be able to calculate the radians for "+formatDate(test.time), func(t *testing.T) {
			received := hoursInRadians(test.time)
			expected := test.angle

			if !approximatedFloat64(received, expected) {
				t.Fatalf("❌ expected %v received %v", expected, received)
			}
		})
	}

}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, test := range cases {
		t.Run("should be able to calculate the point for "+formatDate(test.time), func(t *testing.T) {
			received := secondHandPoint(test.time)
			assertPoint(t, received, test.point)
		})
	}
}

func TestMinuteHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, test := range cases {
		t.Run("should be able to calculate the point for "+formatDate(test.time), func(t *testing.T) {
			received := minuteHandPoint(test.time)
			assertPoint(t, received, test.point)
		})
	}
}

func TestHourHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, test := range cases {
		t.Run("should be able to calculate the point for "+formatDate(test.time), func(t *testing.T) {
			received := hourHandPoint(test.time)
			assertPoint(t, received, test.point)
		})
	}
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(formatDate(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			errSVG := SVGWriter(&b, c.time)

			svg := SVG{}
			err := xml.Unmarshal(b.Bytes(), &svg)

			if err != nil || errSVG != nil {
				t.Errorf("❌ error while parsing XML")
			}

			if !containsLine(c.line, svg.Line) {
				t.Errorf("❌ expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 70},
		},
	}

	for _, c := range cases {
		t.Run(formatDate(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			errSVG := SVGWriter(&b, c.time)

			svg := SVG{}
			err := xml.Unmarshal(b.Bytes(), &svg)

			if err != nil || errSVG != nil {
				t.Errorf("❌ error while parsing XML")
			}

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(6, 0, 0),
			Line{150, 150, 150, 200},
		},
	}

	for _, c := range cases {
		t.Run(formatDate(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			errSVG := SVGWriter(&b, c.time)

			svg := SVG{}
			err := xml.Unmarshal(b.Bytes(), &svg)

			if err != nil || errSVG != nil {
				t.Errorf("❌ error while parsing XML")
			}

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2019, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

func formatDate(t time.Time) string {
	return t.Format("15:04")
}

func approximatedFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func approximatedPoint(a, b Point) bool {
	return approximatedFloat64(a.X, b.X) &&
		approximatedFloat64(a.Y, b.Y)
}

func containsLine(line Line, lines []Line) bool {
	for _, l := range lines {
		if l == line {
			return true
		}
	}
	return false
}

// func TestSecondHandAtMidnight(t *testing.T) {
// 	tm := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)

// 	expected := clockface.Point{X: 150, Y: 150 - 90}
// 	received := clockface.SecondHand(tm)

// 	assertPoint(t, received, expected)
// }

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(2019, time.January, 1, 0, 0, 30, 0, time.UTC)
	buffer := bytes.Buffer{}

	expected := Point{X: 150, Y: 150 + 90}
	secondHand(&buffer, tm)

	fmt.Printf("%v", expected)
}

func assertPoint(t *testing.T, received, expected Point) {
	t.Helper()
	if !approximatedPoint(received, expected) {
		t.Fatalf("❌ received %v expected %v", received, expected)
	}
}
