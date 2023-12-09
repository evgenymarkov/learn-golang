package clockface

import (
	"math"
	"time"
)

type point struct {
	X float64
	Y float64
}

func getSecondHandPoint(t time.Time) point {
	angle := convertSecondsToRadians(t)

	return transformAngleToPoint(angle)
}

func getMinuteHandPoint(t time.Time) point {
	angle := convertMinutesToRadians(t)

	return transformAngleToPoint(angle)
}

func getHourHandPoint(t time.Time) point {
	angle := convertHoursToRadians(t)

	return transformAngleToPoint(angle)
}

func transformAngleToPoint(angle float64) point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return point{x, y}
}

func convertSecondsToRadians(t time.Time) float64 {
	return math.Pi / (30 / (float64(t.Second())))
}

func convertMinutesToRadians(t time.Time) float64 {
	return convertSecondsToRadians(t)/60 + math.Pi/(30/float64(t.Minute()))
}

func convertHoursToRadians(t time.Time) float64 {
	return convertMinutesToRadians(t)/12 + math.Pi/(6/float64(t.Hour()%12))
}
