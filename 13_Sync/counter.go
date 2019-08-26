package sync

import "sync"

// Counter struct with value propriety
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter returns a pointer with a Counter instance
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments the value propriety. (Locks)
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the value
func (c *Counter) Value() int {
	return c.value
}
