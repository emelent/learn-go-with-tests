package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if got := counter.Value(); got != 3 {
			t.Errorf("got %d, want %d", got, 3)
		}
	})
}
