package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	line := input.ReadInputFileToLines("input.txt")[0]

	nums := make(map[string]int)
	for _, num := range strings.Split(line, " ") {
		nums[num] = 1
	}

	for i := 0; i < 75; i++ {
		nums = blink(nums)
		if i == 24 {
			fmt.Printf("Number of stones (25 blinks): %d\n", helpers.SumMap(nums))
		}
	}
	fmt.Printf("Number of stones (75 blinks): %d\n", helpers.SumMap(nums))
}

func blink(nums map[string]int) map[string]int {
	newNums := make(map[string]int)
	for num, count := range nums {
		switch {
		case num == "0":
			newNums["1"] += count
		case len(num)%2 == 0:
			leftNum := num[:len(num)/2]
			rightNum := strings.TrimLeft(num[len(num)/2:], "0")
			if rightNum == "" {
				rightNum = "0"
			}
			newNums[leftNum] += count
			newNums[rightNum] += count
		default:
			mathNum, _ := strconv.Atoi(num)
			newNum := strconv.Itoa(mathNum * 2024)
			newNums[newNum] += count
		}
	}
	return newNums
}
