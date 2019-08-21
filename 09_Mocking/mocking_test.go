package countdown

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const WRITE = "write"
const SLEEP = "sleep"

type CountdownOperationsSpy struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, SLEEP)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, WRITE)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("should starting to count on 3", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		received := buffer.String()
		expected := `3
2
1
GO!`
		assertString(t, received, expected)
	})

	t.Run("should be able to return the correct order of the countdown", func(t *testing.T) {
		spySleeper := &CountdownOperationsSpy{}
		Countdown(spySleeper, spySleeper)

		expected := []string{
			SLEEP,
			WRITE,
			SLEEP,
			WRITE,
			SLEEP,
			WRITE,
			SLEEP,
			WRITE,
		}

		assertSliceSpy(t, spySleeper, expected)
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

func assertString(t *testing.T, received string, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("❌ received %q expected %s", received, expected)
	}
}

func assertSliceSpy(t *testing.T, spy *CountdownOperationsSpy, expected []string) {
	t.Helper()
	if !reflect.DeepEqual(expected, spy.Calls) {
		t.Errorf("❌ wanted calls %v got %v", expected, spy.Calls)
	}
}
