package countdown

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const finalword = "GO!"
const countdownstart = 3

// Sleeper interface with Sleep() function
type Sleeper interface {
	Sleep()
}

// DefaultSleeper default sleeper
type DefaultSleeper struct{}

// Sleep waits for a second
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// ConfigurableSleeper struct with duration and sleep
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep waits for given seconds
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown counts to given time
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownstart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprint(writer, strconv.Itoa(i)+"\n")
	}
	sleeper.Sleep()
	fmt.Fprint(writer, finalword)
}

/*
func main() {
    sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
    Countdown(os.Stdout, sleeper)
}
*/
