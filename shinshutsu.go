package main

import (
	"./bell"
	"./timer"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	teas = map[string][]int{
		"nisemono-cha":        []int{2, 1, 3},
		"long-jing-cha (80C)": []int{20, 35, 70, 120},
		"chun-mei-cha (95C)":  []int{15, 25, 35, 50, 80, 110, 140},
		"chun-mei-cha (80C)":  []int{3 * 60, 3*60 + 20, 3*60 + 40, 4 * 60},
		"zhu-cha (95C)":       []int{10, 25, 40, 50, 70, 90},
		// "four-seasons (80C)":  []int{25, 40, 50, 60, 70, 80, 90, 100},
		// "four-seasons (70C)":  []int{120, 150, 180, 210, 240},
	}
	p    = fmt.Print
	puts = fmt.Println
	keys []string
)

func generateKeys() {
	for k := range teas {
		keys = append(keys, k)
	}
}

func main() {
	generateKeys()
	teaName := keys[bootQuery()]
	steepNumber := getSteepNumber()
	brewingTimes := teas[teaName][(steepNumber - 1):]
	for i, brewingTime := range brewingTimes {
		runTimer(brewingTime)

		// the continue queries should stop if the current brewing time is the last one
		if i == (len(brewingTimes) - 1) {
			puts("\nterminating as there are no brewing times for next steep")
			return
		}

		if userQuits(brewingTime) {
			return
		}
	}
}

func bootQuery() (teaNumber int) {
	for index, teaName := range keys {
		fmt.Printf("[%v] %v\n", index, teaName)
	}
	puts()
	p("please choose tea (by number): ")
	reader := bufio.NewReader(os.Stdin)
	teaQueryResponse, _ := reader.ReadByte()
	teaNumber, err := strconv.Atoi(string(teaQueryResponse))
	if err != nil {
		abort(err.Error())
	}
	if teaNumber < 0 || teaNumber >= len(teas) {
		abort("no such number in the list")
	}
	return
}

func abort(msg string) {
	fmt.Fprintf(os.Stderr, "error: %v\n", msg)
	os.Exit(1)
}

func userQuits(brewingTime int) bool {
	p("[c]ontinue/[s]top: ")
	continueResponse, _, _ := bufio.NewReader(os.Stdin).ReadRune()
	switch continueResponse {
	case 'c':
		return false
	case 's':
		return true
	default:
		puts("please enter a valid option")
		return userQuits(brewingTime)
	}
}

func getSteepNumber() (steepNumber int) {
	p("please choose steep number: ")
	reader := bufio.NewReader(os.Stdin)
	steepNumberQueryResponse, _, _ := reader.ReadRune()
	steepNumber, _ = strconv.Atoi(string(steepNumberQueryResponse))
	return
}

func runTimer(seconds int) {
	timer.Timer{seconds}.Countdown()
	bell.Toll()
}
