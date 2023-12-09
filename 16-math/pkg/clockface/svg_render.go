package clockface

import (
	"fmt"
	"io"
	"strings"
	"time"
)

const (
	clockCenterX = 150
	clockCenterY = 150
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
)

// Writes an SVG representation of an analogue clock, showing the time `t`, to the writer `w`
func RenderClock(w io.StringWriter, t time.Time) {
	secondHand := placeSecondHand(t)
	minuteHand := placeMinuteHand(t)
	hourHand := placeHourHand(t)

	w.WriteString(svgStart)
	w.WriteString(svgBezel)
	renderSecondHand(w, secondHand)
	renderMinuteHand(w, minuteHand)
	renderHourHand(w, hourHand)
	w.WriteString(svgEnd)
}

func placeSecondHand(t time.Time) point {
	p := getSecondHandPoint(t)
	p = placeHand(p, secondHandLength)

	return p
}

func placeMinuteHand(t time.Time) point {
	p := getMinuteHandPoint(t)
	p = placeHand(p, minuteHandLength)

	return p
}

func placeHourHand(t time.Time) point {
	p := getHourHandPoint(t)
	p = placeHand(p, hourHandLength)

	return p
}

func placeHand(p point, length float64) point {
	p = point{p.X * length, p.Y * length}             // scale
	p = point{p.X, -p.Y}                              // flip
	p = point{p.X + clockCenterX, p.Y + clockCenterY} // translate

	return p
}

func renderSecondHand(w io.StringWriter, p point) {
	w.WriteString(fmt.Sprintf(svgSecondHand, p.X, p.Y))
}

func renderMinuteHand(w io.StringWriter, p point) {
	w.WriteString(fmt.Sprintf(svgMinuteHand, p.X, p.Y))
}

func renderHourHand(w io.StringWriter, p point) {
	w.WriteString(fmt.Sprintf(svgHourHand, p.X, p.Y))
}

var svgStart = strings.TrimSpace(`
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg" width="100%" height="100%" viewBox="0 0 300 300" version="2.0">
`)

var svgBezel = strings.TrimSpace(`
<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>
`)

var svgEnd = strings.TrimSpace(`
</svg>
`)

var svgSecondHand = strings.TrimSpace(`
<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>
`)

var svgMinuteHand = strings.TrimSpace(`
<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>
`)

var svgHourHand = strings.TrimSpace(`
<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>
`)
