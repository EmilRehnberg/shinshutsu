package main

import (
	"fmt"
	"github.com/emilrehnberg/shinshutsu/alarmclock"
	"os"
	"strconv"
)

func main() {
	alarmclock.Timer{parseSecondsFromArgs()}.Countdown()
	alarmclock.Toll()
}

func parseSecondsFromArgs() (seconds int) {
	puts := fmt.Println
	args := os.Args
	if len(args) != 2 {
		puts("please call with the amount of seconds. E.g.")
		puts("> countdown 20")
		return 0
	}
	seconds, _ = strconv.Atoi(args[1])
	return
}
