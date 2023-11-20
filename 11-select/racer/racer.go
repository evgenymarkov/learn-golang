package racer

import (
	"fmt"
	"net/http"
	"time"
)

const defaultRaceTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, defaultRaceTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	channel := make(chan struct{})

	go func() {
		http.Get(url)
		close(channel)
	}()

	return channel
}
