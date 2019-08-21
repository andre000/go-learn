package countdown

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const FINAL_WORD = "GO!"
const COUNTDOWN_START = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := COUNTDOWN_START; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprint(writer, strconv.Itoa(i)+"\n")
	}
	sleeper.Sleep()
	fmt.Fprint(writer, FINAL_WORD)
}

/*
func main() {
    sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
    Countdown(os.Stdout, sleeper)
}
*/
