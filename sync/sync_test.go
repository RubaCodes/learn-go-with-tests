package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		// A WaitGroup waits for a collection of goroutines to finish.
		// The main goroutine calls Add to set the number of goroutines to wait for.
		// Then each of the goroutines runs and calls Done when finished.
		// At the same time, Wait can be used to block until all goroutines have finished.

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for range wantedCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assetCounter(t, counter, wantedCount)
	})
}

func assetCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
func NewCounter() *Counter {
	return &Counter{}
}

// Paraphrasing:

// Use channels when passing ownership of data

// Use mutexes for managing state
