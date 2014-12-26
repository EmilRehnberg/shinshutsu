package main

import (
	"./bell"
	"./timer"
	"fmt"
	"os"
	"strconv"
)

func main() {
	timer.Timer{parseSecondsFromArgs()}.Countdown()
	bell.Toll()
}

func parseSecondsFromArgs() (seconds int) {
	puts := fmt.Println
	args := os.Args
	if len(args) != 2 {
		puts("please call with the amount of seconds. E.g.")
		puts("> countdown-timer 20")
		return 0
	}
	seconds, _ = strconv.Atoi(args[1])
	return
}
