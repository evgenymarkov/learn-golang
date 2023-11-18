package countdown

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	sleepOperation = "sleep"
	writeOperation = "write"
)

type SpyTime struct {
	durationSlept time.Duration
}

func (t *SpyTime) Sleep(d time.Duration) {
	t.durationSlept += d
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleepOperation)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, writeOperation)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go", func(t *testing.T) {
		printer := &bytes.Buffer{}
		sleeper := &SpyCountdownOperations{}

		Countdown(printer, sleeper)

		got := printer.String()
		want := "3\n2\n1\nGo"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("sleeps before every print", func(t *testing.T) {
		spyOperations := &SpyCountdownOperations{}

		Countdown(spyOperations, spyOperations)

		want := []string{
			writeOperation, // 3
			sleepOperation,
			writeOperation, // 2
			sleepOperation,
			writeOperation, // 1
			sleepOperation,
			writeOperation, // Go
		}

		if !reflect.DeepEqual(spyOperations.Calls, want) {
			t.Errorf("want calls %v, got calls %v", want, spyOperations.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	spyTime := &SpyTime{}
	sleepTime := 5 * time.Second

	sleeper := &ConfigurableSleeper{
		Duration:      sleepTime,
		SleepFunction: spyTime.Sleep,
	}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
