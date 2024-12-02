package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	totalSafe := 0
	totalSafeDampened := 0

	for _, line := range lines {
		numStrings := strings.Split(line, " ")
		nums := make([]float64, len(numStrings))
		for i, numString := range numStrings {
			num, err := strconv.Atoi(numString)
			helpers.Check(err)
			nums[i] = float64(num)
		}

		if evalNums(nums) {
			totalSafe++
			totalSafeDampened++
		} else {
			// Remove one of the numbers and re-evaluate, for every number
			for i := 0; i < len(nums); i++ {
				newNums := make([]float64, len(nums)-1)
				copy(newNums, nums[:i])
				copy(newNums[i:], nums[i+1:])
				if evalNums(newNums) {
					totalSafeDampened++
					break
				}
			}
		}
	}

	fmt.Printf("Safe: %d\n", totalSafe)
	fmt.Printf("Safe (dampened): %d\n", totalSafeDampened)
}

func evalNums(nums []float64) bool {
	decreasing := nums[0] > nums[1]
	for i := 0; i < len(nums)-1; i++ {
		absDiff := math.Abs(nums[i] - nums[i+1])
		localDecreasing := nums[i] > nums[i+1]
		if localDecreasing != decreasing || absDiff < 1.0 || absDiff > 3.0 {
			return false
		}
	}
	return true
}
