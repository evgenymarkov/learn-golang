package counter

import "sync"

type Counter struct {
	mu    sync.Mutex
	calls int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.calls++
}

func (c *Counter) Value() int {
	return c.calls
}
