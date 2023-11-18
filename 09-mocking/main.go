package main

import (
	"os"
	"time"

	"github.com/evgenymarkov/learn-golang/09-mocking/countdown"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{
		Duration:      1 * time.Second,
		SleepFunction: time.Sleep,
	}
	countdown.Countdown(os.Stdout, sleeper)
}
