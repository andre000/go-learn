package selectgo

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	aDuration, errA := measureResponseTime(a)

	if errA != nil {
		return "", errA
	}

	bDuration, errB := measureResponseTime(b)

	if errB != nil {
		return "", errB
	}

	if aDuration < bDuration {
		return a, nil
	}

	return b, nil
}

func measureResponseTime(url string) (time.Duration, error) {
	start := time.Now()
	_, err := http.Get(url)
	return time.Since(start), err
}
