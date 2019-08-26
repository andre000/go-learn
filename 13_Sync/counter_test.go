package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("should increment the counter 3 times returning 3 at the end", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		expected := 3
		assertCounter(t, counter, expected)
	})

	t.Run("should be able to run concurrently", func(t *testing.T) {
		expected := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(expected)

		for i := 0; i < expected; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}

		wg.Wait()
		assertCounter(t, counter, expected)
	})
}

func assertCounter(t *testing.T, received *Counter, expected int) {
	t.Helper()
	if received.Value() != expected {
		t.Errorf("❌ received %d expected %d", received.Value(), expected)
	}
}
