package brewci

import (
	"bufio"
	"fmt"
	"github.com/EmilRehnberg/shinshutsu/alarmclock"
	"os"
	"strconv"
	"strings"
)

var (
	teaNames []string
	p        = fmt.Print
	printf   = fmt.Printf
	puts     = fmt.Println
)

// executes the command line user interaction (for the brewing times given).
func Execute(teas map[string][]int) {
	brewingTimes := getBrewingTimes(teas)
	for i, brewingTime := range brewingTimes {
		runTimer(brewingTime)

		// the continue queries should stop if the current brewing time is the last one
		if i == (len(brewingTimes) - 1) {
			impartUserOnMissingNextSteep()
			return
		}

		if userQuits() {
			return
		}
	}
}

func getBrewingTimes(teas map[string][]int) []int {
	return teaFullBrewingTimes(teas)[(getSteepNumber() - 1):]
}

func teaFullBrewingTimes(teas map[string][]int) []int {
	return teas[teaQuery(teas)]
}

func teaQuery(teas map[string][]int) string {
	setTeaNames(teas)
	printTeaSelection(teaNames)
	return teaNames[getTeaNumber(teas)]
}

func printTeaSelection(teaNames []string) {
	for index, teaName := range teaNames {
		printf("[%v] %v\n", index, teaName)
	}
	puts()
}

func getTeaNumber(teas map[string][]int) int {
	p("please choose tea (by number): ")
	return queryForTeaNumber(teas)
}

func queryForTeaNumber(teas map[string][]int) (teaNumber int) {
	teaNumber = parseInt(teaQueryResponse())
	if teaNumber < 0 || teaNumber >= len(teas) {
		abort("no such number in the list")
	}
	return
}

var teaQueryResponse = func() string {
	return readUserString()
}

func userQuits() bool {
	p("[c]ontinue/(s)top: ")
	switch continueResponse() {
	case "c", "":
		return false
	case "s":
		return true
	default:
		puts("please enter a valid option")
		return userQuits()
	}
}

var continueResponse = func() string {
	return strings.TrimSpace(readUserString())
}

func getSteepNumber() int {
	p("please choose steep number [1]: ")
	return parseInt(steepNumber())
}

func steepNumber() string {
	requestedSteepNumber := steepNumberQueryResponse()
	if requestedSteepNumber == "\n" {
		return "1"
	}
	return requestedSteepNumber
}

var steepNumberQueryResponse = func() string {
	return readUserString()
}

func impartUserOnMissingNextSteep() {
	puts("terminating as there are no brewing times for next steep")
	return
}

func readUserString() string {
	readString, err := reader().ReadString('\n')
	handleError(err)
	return readString
}

func abort(msg string) {
	fmt.Fprintf(os.Stderr, "error: %v\n", msg)
	os.Exit(1)
}

func setTeaNames(teas map[string][]int) {
	for k := range teas {
		teaNames = append(teaNames, k)
	}
}

var runTimer = func(seconds int) {
	alarmclock.Timer{seconds}.Countdown()
	alarmclock.Toll()
}

func handleError(err error) {
	if err != nil {
		abort(err.Error())
	}
	return
}

func reader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func parseInt(unparsed string) (parsed int) {
	parsed, err := strconv.Atoi(strings.TrimSpace(unparsed))
	handleError(err)
	return
}
