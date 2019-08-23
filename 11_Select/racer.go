package selectgo

import (
	"errors"
	"net/http"
	"time"
)

// ErrTimeout returned when request takes more than the defined timeout
var ErrTimeout = errors.New("❌ Error! Server timeout")
var defaultTimeout = 10 * time.Second

// ConfigurableRacer Racer with configurable timeout
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", ErrTimeout
	}
}

// Racer takes two urls and returns the fastest one
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, defaultTimeout)
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		_, err := http.Get(url)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}()
	return ch
}
