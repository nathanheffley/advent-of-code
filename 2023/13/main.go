package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	groups := [][]string{}
	currentGroup := []string{}
	for _, line := range lines {
		if line == "" {
			groups = append(groups, currentGroup)
			currentGroup = []string{}
		} else {
			currentGroup = append(currentGroup, line)
		}
	}
	if len(currentGroup) > 0 {
		groups = append(groups, currentGroup)
	}

	part1Total := 0
	for _, group := range groups {
		isReflection := false

		// Find vertical reflections
		for i := 0; i < len(group)-1; i++ {
			if group[i] == group[i+1] {
				ii := i + 2
				isReflection = true
				for j := i - 1; j >= 0 && ii < len(group) && isReflection; j-- {
					if group[j] != group[ii] {
						isReflection = false
					}
					ii++
				}
				if isReflection {
					part1Total += (i + 1) * 100
					break
				}
			}
		}
		if isReflection {
			continue
		}

		// Find horizontal reflections
		columnGroup := make([]string, len(group[0]))
		for _, line := range group {
			for i, char := range line {
				columnGroup[i] = columnGroup[i] + string(char)
			}
		}
		for i := 0; i < len(columnGroup)-1; i++ {
			if columnGroup[i] == columnGroup[i+1] {
				ii := i + 2
				isReflection = true
				for j := i - 1; j >= 0 && ii < len(columnGroup) && isReflection; j-- {
					if columnGroup[j] != columnGroup[ii] {
						isReflection = false
					}
					ii++
				}
				if isReflection {
					part1Total += i + 1
					break
				}
			}
		}
	}

	fmt.Println(part1Total)
}
