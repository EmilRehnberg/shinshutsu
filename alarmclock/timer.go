package alarmclock

import (
	"fmt"
	"time"
)

var (
	tprintf = fmt.Printf
)

type Timer struct {
	Seconds int
}

var newTicker = func() ClockTicker {
	return &timeTicker{time.NewTicker(second())}
}

// Runs a countdown for the seconds on the Timer object sent to.
func (clock Timer) Countdown() {
	ticker := newTicker()
	clock.runCountdown(ticker)
}

func (clock Timer) runCountdown(ticker ClockTicker) {
	for remaining := clock.duration(); remaining >= 0; remaining -= second() {
		tprintf(formattedDuration(remaining))
		ticker.Pop()
	}
	ticker.Stop()
}

func (clock Timer) duration() time.Duration {
	return second() * time.Duration(clock.Seconds)
}

type ClockTicker interface {
	Stop()
	Pop()
}

type timeTicker struct {
	time *time.Ticker
}

func (t *timeTicker) Stop() {
	t.time.Stop()
}

func (t *timeTicker) Pop() {
	<-t.time.C
}

func formattedDuration(duration time.Duration) string {
	return fmt.Sprintf(
		"%02d:%02d%15s\r",
		minutesRemaining(duration),
		secondsRemaining(duration),
		" ",
	)
}

func minutesRemaining(timeLeft time.Duration) int {
	return int(timeLeft.Minutes())
}

func secondsRemaining(timeLeft time.Duration) int {
	return int(timeLeft.Seconds()) - (int(timeLeft.Minutes()) * 60)
}

func second() time.Duration {
	return time.Second
}
