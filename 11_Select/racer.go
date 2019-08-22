package selectgo

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	startA := time.Now()
	_, errA := http.Get(a)
	aDuration := time.Since(startA)

	if errA != nil {
		return "", errA
	}

	startB := time.Now()
	_, errB := http.Get(b)
	bDuration := time.Since(startB)

	if errB != nil {
		return "", errB
	}

	if aDuration < bDuration {
		return a, nil
	}

	return b, nil
}
