package main

import (
	"github.com/emilrehnberg/shinshutsu/brewci"
)

var (
	teas = map[string][]int{
		"nisemono-cha":        []int{2, 1, 3},
		"long-jing-cha (80C)": []int{20, 35, 70, 120},
		"chun-mei-cha (80C)":  []int{120, 150, 180, 210},
		"nihon-sencha (70C)":  []int{180, 200, 220},
		"nihon-sencha (80C)":  []int{20, 35, 70, 120},
		"zhu-cha (95C)":       []int{10, 25, 40, 50, 70, 90},
	}
)

func main() {
	brewci.Execute(teas)
}
