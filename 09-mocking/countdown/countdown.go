package countdown

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration      time.Duration
	SleepFunction func(d time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.SleepFunction(s.Duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
