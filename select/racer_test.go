package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const slowServerDelay = 20 * time.Millisecond
const fastServerDelay = 0 * time.Millisecond
const timeoutDelay = 40 * time.Millisecond

func TestRacer(t *testing.T) {
	t.Run("resturns fastest url", func(t *testing.T) {
		slowServer := makeDelayedServer(slowServerDelay)
		fastServer := makeDelayedServer(fastServerDelay)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		slowServer.Close()
		fastServer.Close()
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(60 * time.Millisecond)
		serverB := makeDelayedServer(60 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, timeoutDelay)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func BenchmarkSequentialRacer(b *testing.B) {
	slowServer := makeDelayedServer(slowServerDelay)
	defer slowServer.Close()
	fastServer := makeDelayedServer(fastServerDelay)
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	for i := 0; i < b.N; i++ {
		SequentialRacer(slowURL, fastURL)
	}
}

func BenchmarkParallelRacer(b *testing.B) {
	slowServer := makeDelayedServer(slowServerDelay)
	defer slowServer.Close()
	fastServer := makeDelayedServer(fastServerDelay)
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	for i := 0; i < b.N; i++ {
		ParallelRacer(slowURL, fastURL)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		},
	))
}
