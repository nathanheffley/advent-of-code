package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	var ranges [][2]int
	var availableIngredients []int

	findingRanges := true
	for _, line := range lines {
		if line == "" {
			findingRanges = false
			continue
		}

		if findingRanges {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				continue
			}

			start, err := strconv.Atoi(parts[0])
			helpers.Check(err)
			end, err := strconv.Atoi(parts[1])
			helpers.Check(err)

			ranges = append(ranges, [2]int{start, end})
		} else {
			ingredient, err := strconv.Atoi(line)
			helpers.Check(err)
			availableIngredients = append(availableIngredients, ingredient)
		}
	}

	ranges = simplifyRanges(ranges)

	freshAvailableIngredientsCount := 0
	for _, ingredient := range availableIngredients {
		for _, r := range ranges {
			if ingredient >= r[0] && ingredient <= r[1] {
				freshAvailableIngredientsCount++
				break
			}
		}
	}

	totalFreshIngredients := 0
	for _, r := range ranges {
		totalFreshIngredients += r[1] - r[0] + 1
	}

	fmt.Printf("Fresh available ingredients: %d\n", freshAvailableIngredientsCount)
	fmt.Printf("Total fresh ingredients: %d\n", totalFreshIngredients)
}

func simplifyRanges(ranges [][2]int) [][2]int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	for i := 0; i < len(ranges)-1; i++ {
		if ranges[i][1] >= ranges[i+1][0] {
			newRange := [2]int{
				ranges[i][0],
				int(math.Max(float64(ranges[i][1]), float64(ranges[i+1][1]))),
			}

			newRanges := ranges[:i]
			newRanges = append(newRanges, newRange)
			newRanges = append(newRanges, ranges[i+2:]...)
			ranges = simplifyRanges(newRanges)
		}
	}

	return ranges
}
