package main

import (
	"fmt"
	"strconv"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	nums := make([]int, len(lines))
	for i, line := range lines {
		ordinal := line[:1]
		num, _ := strconv.Atoi(line[1:])
		if ordinal == "+" {
			nums[i] = num
		} else {
			nums[i] = -num
		}
	}

	freq := 0
	seenFreqs := make(map[int]bool)
	seenFreqs[0] = true
	var firstRepeatFreq int
	for i := 0; i < len(nums); i++ {
		freq += nums[i]
		if _, ok := seenFreqs[freq]; ok {
			firstRepeatFreq = freq
			break
		} else {
			seenFreqs[freq] = true
		}
		if i == len(nums)-1 {
			i = -1
		}
	}

	fmt.Printf("Frequency: %d\n", helpers.SumSlice(nums))
	fmt.Printf("First Repeated Frequency: %d\n", firstRepeatFreq)
}
