package clockface

import (
	"math"
	"testing"
	"time"
)

func TestGetSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point point
	}{
		{getDayTime(0, 0, 30), point{0, -1}},
		{getDayTime(0, 0, 45), point{-1, 0}},
	}

	for _, test := range cases {
		t.Run(getTestName(test.time), func(t *testing.T) {
			got := getSecondHandPoint(test.time)

			if !isPointsRoughlyEqual(got, test.point) {
				t.Fatalf("got %v, want %v point", got, test.point)
			}
		})
	}
}

func TestGetMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point point
	}{
		{getDayTime(0, 30, 0), point{0, -1}},
		{getDayTime(0, 45, 0), point{-1, 0}},
	}

	for _, test := range cases {
		t.Run(getTestName(test.time), func(t *testing.T) {
			got := getMinuteHandPoint(test.time)

			if !isPointsRoughlyEqual(got, test.point) {
				t.Fatalf("got %v, want %v point", got, test.point)
			}
		})
	}
}

func TestGetHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point point
	}{
		{getDayTime(6, 0, 0), point{0, -1}},
		{getDayTime(21, 0, 0), point{-1, 0}},
	}

	for _, test := range cases {
		t.Run(getTestName(test.time), func(t *testing.T) {
			got := getHourHandPoint(test.time)

			if !isPointsRoughlyEqual(got, test.point) {
				t.Fatalf("got %v, want %v point", got, test.point)
			}
		})
	}
}

func TestConvertSecondsToRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{getDayTime(0, 0, 0), 0},
		{getDayTime(0, 0, 7), (math.Pi / 30) * 7},
		{getDayTime(0, 0, 30), math.Pi},
		{getDayTime(0, 0, 45), (math.Pi / 2) * 3},
	}

	for _, test := range cases {
		t.Run(getTestName(test.time), func(t *testing.T) {
			want := test.angle
			got := convertSecondsToRadians(test.time)

			if !isFloat64RoughlyEqual(got, want) {
				t.Errorf("got %v, want %v radians", got, want)
			}
		})
	}
}

func TestConvertMinutesToRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{getDayTime(0, 30, 0), math.Pi},
		{getDayTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, test := range cases {
		t.Run(getTestName(test.time), func(t *testing.T) {
			want := test.angle
			got := convertMinutesToRadians(test.time)

			if !isFloat64RoughlyEqual(got, want) {
				t.Errorf("got %v, want %v radians", got, want)
			}
		})
	}
}

func TestConvertHoursToRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{getDayTime(6, 0, 0), math.Pi},
		{getDayTime(0, 0, 0), 0},
		{getDayTime(21, 0, 0), math.Pi * 1.5},
		{getDayTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, test := range cases {
		t.Run(getTestName(test.time), func(t *testing.T) {
			want := test.angle
			got := convertHoursToRadians(test.time)

			if !isFloat64RoughlyEqual(got, want) {
				t.Errorf("got %v, want %v radians", got, want)
			}
		})
	}
}

func getTestName(t time.Time) string {
	return t.Format("15:04:05")
}

func getDayTime(h, m, s int) time.Time {
	return time.Date(1970, time.January, 0, h, m, s, 0, time.UTC)
}

func isPointsRoughlyEqual(a, b point) bool {
	return isFloat64RoughlyEqual(a.X, b.X) && isFloat64RoughlyEqual(a.Y, b.Y)
}

func isFloat64RoughlyEqual(a, b float64) bool {
	const equalityThreshold = 1e-7

	return math.Abs(a-b) < equalityThreshold
}
