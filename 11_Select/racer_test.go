package selectgo

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

func TestRacer(t *testing.T) {
	t.Run("should be able to return the fastest url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expected := fastURL
		received, err := Racer(slowURL, fastURL)

		assertNoError(t, err)
		assertString(t, received, expected)
	})

	t.Run("should return an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)
		assertError(t, err, ErrTimeout)
	})
}

func assertString(t *testing.T, received string, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("❌ received %q expected %q", received, expected)
	}
}

func assertNoError(t *testing.T, received error) {
	t.Helper()
	if received != nil {
		t.Error("❌ received error expected nil")
	}
}

func assertError(t *testing.T, received error, expected error) {
	t.Helper()
	if received != expected {
		t.Errorf("❌ expected error %q received %q", received, expected)
	}
}
