package whos

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestWhosFasterURL(t *testing.T) {
	t.Run("choose fastest url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(12 * time.Millisecond)
		ultraServer := makeDelayedServer(3 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		defer ultraServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		ultraURL := ultraServer.URL

		want := ultraURL
		got, err := WhosFasterURL(slowURL, ultraURL, fastURL)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns an error if a server doesn't respond within 25ms", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)
		defer server.Close()
		_, err := ConfigurableWhosFasterURL(20*time.Millisecond, server.URL, server.URL, server.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
