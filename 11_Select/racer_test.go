package selectgo

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	expected := fastURL
	received, err := Racer(slowURL, fastURL)

	assertNoError(t, err)
	assertString(t, received, expected)

	slowServer.Close()
	fastServer.Close()
}

func assertString(t *testing.T, received string, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("❌ received %q expected %q", received, expected)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error("❌ received error expected nil")
	}
}
