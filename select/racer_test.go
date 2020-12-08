package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const slowServerDelay = 20
const fastServerDelay = 0

func TestRacer(t *testing.T) {
	t.Run("resturns fastest url", func(t *testing.T) {
		slowServer := makeDelayedServer(slowServerDelay * time.Millisecond)
		fastServer := makeDelayedServer(fastServerDelay * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		slowServer.Close()
		fastServer.Close()
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func BenchmarkSequentialRacer(b *testing.B) {
	slowServer := makeDelayedServer(slowServerDelay * time.Millisecond)
	defer slowServer.Close()
	fastServer := makeDelayedServer(fastServerDelay * time.Millisecond)
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	for i := 0; i < b.N; i++ {
		SequentialRacer(slowURL, fastURL)
	}
}

func BenchmarkParallelRacer(b *testing.B) {
	slowServer := makeDelayedServer(slowServerDelay * time.Millisecond)
	defer slowServer.Close()
	fastServer := makeDelayedServer(fastServerDelay * time.Millisecond)
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
