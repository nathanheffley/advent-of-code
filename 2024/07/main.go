package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	start := time.Now()

	validTotals := 0
	part2ValidTotals := 0
	for _, line := range lines {
		numRegex := regexp.MustCompile(`\d+`)
		numStrings := numRegex.FindAllString(line, -1)
		nums := make([]int, 0)
		for _, numString := range numStrings {
			num, _ := strconv.Atoi(numString)
			nums = append(nums, num)
		}
		part1Valid, part2Valid := valid(nums)
		if part1Valid {
			validTotals += nums[0]
		}
		if part2Valid {
			part2ValidTotals += nums[0]
		}
	}

	fmt.Printf("Part 1: %d\n", validTotals)
	fmt.Printf("Part 2: %d\n", part2ValidTotals)

	elapsed := time.Since(start)
	fmt.Printf("(took %s)\n", elapsed)
}

func valid(nums []int) (bool, bool) {
	results := make([]int, 0)
	part2Results := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		if len(results) == 0 {
			results = append(results, nums[i])
			part2Results = append(results, nums[i])
			continue
		}

		newResults := make([]int, 0)
		for j := 0; j < len(results); j++ {
			newResults = append(newResults, results[j]+nums[i])
			newResults = append(newResults, results[j]*nums[i])
		}
		results = newResults

		newPart2Results := make([]int, 0)
		for j := 0; j < len(part2Results); j++ {
			newPart2Results = append(newPart2Results, part2Results[j]+nums[i])

			mul := part2Results[j] * nums[i]
			if mul <= nums[0] {
				newPart2Results = append(newPart2Results, mul)
			}

			concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", part2Results[j], nums[i]))
			if concatenated <= nums[0] {
				newPart2Results = append(newPart2Results, concatenated)
			}
		}
		part2Results = newPart2Results
	}

	valid := false
	for _, result := range results {
		if result == nums[0] {
			valid = true
			break
		}
	}

	part2Valid := false
	for _, result := range part2Results {
		if result == nums[0] {
			part2Valid = true
			break
		}
	}

	return valid, part2Valid
}
