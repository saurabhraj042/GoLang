package sync

import(
	"testing"
	"sync"
)

func TestSync(t *testing.T) {
	t.Run("Run three times to get value: 3", func(t *testing.T) {
		counter := NewCounter()

		for i := 0; i < 3; i++ {
			counter.Inc()
		}

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		counter := NewCounter()
		wantedCount := 1000

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value != want {
		t.Errorf("got %d want %d", got.Value, want)
	}
}