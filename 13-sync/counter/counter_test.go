package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		counter := NewCounter()
		wantedCount := 1000

		var wgCounter sync.WaitGroup
		wgCounter.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wgCounter.Done()
			}()
		}
		wgCounter.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, counter *Counter, expectedValue int) {
	t.Helper()
	gotValue := counter.Value()
	if gotValue != expectedValue {
		t.Errorf("got %d, want %d", gotValue, expectedValue)
	}
}
