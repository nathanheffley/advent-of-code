package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	fullOverlap := 0
	partialOverlap := 0

	for _, line := range lines {
		first := strings.Split(line, ",")[0]
		firstLow, _ := strconv.Atoi(string(strings.Split(first, "-")[0]))
		firstHigh, _ := strconv.Atoi(string(strings.Split(first, "-")[1]))

		second := strings.Split(line, ",")[1]
		secondLow, _ := strconv.Atoi(string(strings.Split(second, "-")[0]))
		secondHigh, _ := strconv.Atoi(string(strings.Split(second, "-")[1]))

		if firstLow >= secondLow && firstHigh <= secondHigh {
			fullOverlap++
			partialOverlap++
		} else if secondLow >= firstLow && secondHigh <= firstHigh {
			fullOverlap++
			partialOverlap++
		} else if firstLow <= secondHigh && firstHigh >= secondHigh {
			partialOverlap++
		} else if firstLow <= secondLow && firstHigh >= secondLow {
			partialOverlap++
		}
	}

	fmt.Println(fullOverlap)
	fmt.Println(partialOverlap)
}
