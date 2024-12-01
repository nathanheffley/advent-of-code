package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	leftNums := make([]int, 0)
	rightNums := make([]int, 0)
	for _, line := range lines {
		nums := strings.Split(line, "   ")

		leftNum, err := strconv.Atoi(nums[0])
		helpers.Check(err)
		leftNums = append(leftNums, leftNum)

		rightNum, err := strconv.Atoi(nums[1])
		helpers.Check(err)
		rightNums = append(rightNums, rightNum)
	}

	sort.Slice(leftNums, func(i, j int) bool {
		return leftNums[i] < leftNums[j]
	})

	sort.Slice(rightNums, func(i, j int) bool {
		return rightNums[i] < rightNums[j]
	})

	difference := 0
	similarity := 0

	for i, leftNum := range leftNums {
		// Difference
		rightNumForDiff := rightNums[i]
		if leftNum > rightNumForDiff {
			difference += leftNum - rightNumForDiff
		} else {
			difference += rightNumForDiff - leftNum
		}

		// Similarity
		rightNumCount := 0
		for _, rightNum := range rightNums {
			if rightNum == leftNum {
				rightNumCount += 1
			}
		}
		similarity += leftNum * rightNumCount
	}

	fmt.Printf("Difference: %d\n", difference)
	fmt.Printf("Similarity: %d\n", similarity)
}
