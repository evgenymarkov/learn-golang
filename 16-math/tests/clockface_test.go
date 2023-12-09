package tests

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/evgenymarkov/learn-golang/16-math/pkg/clockface"
)

type clockSVG struct {
	XMLName xml.Name       `xml:"svg"`
	Text    string         `xml:",chardata"`
	Xmlns   string         `xml:"xmlns,attr"`
	Width   string         `xml:"width,attr"`
	Height  string         `xml:"height,attr"`
	ViewBox string         `xml:"viewBox,attr"`
	Version string         `xml:"version,attr"`
	Circle  clockCircleSVG `xml:"circle"`
	Line    []clockLineSVG `xml:"line"`
}

type clockCircleSVG struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type clockLineSVG struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestRenderSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line clockLineSVG
	}{
		{
			time: getDayTime(0, 0, 0),
			line: clockLineSVG{150, 150, 150, 60},
		},
		{
			time: getDayTime(0, 0, 30),
			line: clockLineSVG{150, 150, 150, 240},
		},
	}

	for _, test := range cases {
		t.Run(getTestName(test.time), func(t *testing.T) {
			buffer := bytes.Buffer{}
			clockface.RenderClock(&buffer, test.time)

			gotClockSVG := clockSVG{}
			xml.Unmarshal(buffer.Bytes(), &gotClockSVG)

			for _, line := range gotClockSVG.Line {
				if line == test.line {
					return
				}
			}

			t.Errorf(
				"expected to find the second hand line %+v, in the SVG lines %+v",
				test.line,
				gotClockSVG.Line,
			)
		})
	}
}

func TestRenderMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line clockLineSVG
	}{
		{
			time: getDayTime(0, 0, 0),
			line: clockLineSVG{150, 150, 150, 70},
		},
	}

	for _, test := range cases {
		t.Run(getTestName(test.time), func(t *testing.T) {
			buffer := bytes.Buffer{}
			clockface.RenderClock(&buffer, test.time)

			gotClockSVG := clockSVG{}
			xml.Unmarshal(buffer.Bytes(), &gotClockSVG)

			for _, line := range gotClockSVG.Line {
				if line == test.line {
					return
				}
			}

			t.Errorf(
				"expected to find the minute hand line %+v, in the SVG lines %+v",
				test.line,
				gotClockSVG.Line,
			)
		})
	}
}

func TestRenderHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line clockLineSVG
	}{
		{
			time: getDayTime(6, 0, 0),
			line: clockLineSVG{150, 150, 150, 200},
		},
	}

	for _, test := range cases {
		t.Run(getTestName(test.time), func(t *testing.T) {
			buffer := bytes.Buffer{}
			clockface.RenderClock(&buffer, test.time)

			gotClockSVG := clockSVG{}
			xml.Unmarshal(buffer.Bytes(), &gotClockSVG)

			for _, line := range gotClockSVG.Line {
				if line == test.line {
					return
				}
			}

			t.Errorf(
				"expected to find the hour hand line %+v, in the SVG lines %+v",
				test.line,
				gotClockSVG.Line,
			)
		})
	}
}

func getTestName(t time.Time) string {
	return t.Format("15:04:05")
}

func getDayTime(h, m, s int) time.Time {
	return time.Date(1970, time.January, 0, h, m, s, 0, time.UTC)
}
