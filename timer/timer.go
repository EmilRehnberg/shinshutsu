package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	Seconds int
}

func (clock Timer) Countdown() {
	interval := time.Second
	duration := time.Second * time.Duration(clock.Seconds)
	ticker := time.NewTicker(interval)
	for remaining := duration; remaining >= 0; remaining -= interval {
		printFormattedDuration(remaining)
		<-ticker.C
	}
	ticker.Stop()
}

func minutesRemaining(timeLeft time.Duration) int {
	return int(timeLeft.Minutes())
}

func secondsRemaining(timeLeft time.Duration) int {
	return int(timeLeft.Seconds()) - (int(timeLeft.Minutes()) * 60)
}

func printFormattedDuration(duration time.Duration) {
	fmt.Printf("%02d:%02d      \r",
		minutesRemaining(duration),
		secondsRemaining(duration))
}
