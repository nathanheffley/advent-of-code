package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	total := 0
	c := make(chan int)
	for _, line := range lines {
		go ProcessLine(line, c)
	}
	waitingCount := len(lines)
	for waitingCount > 0 {
		total += <-c
		fmt.Println(total)
		waitingCount--
	}

	fmt.Printf("Part 1 Total: %d\n", total)
}

func ProcessLine(line string, ch chan int) {
	newLines := []string{""}

	data := strings.Split(line, " ")[0]
	data = fmt.Sprintf("%s?%s?%s?%s?%s", data, data, data, data, data)

	checksLine := strings.Split(line, " ")[1]
	checksLine = fmt.Sprintf("%s,%s,%s,%s,%s", checksLine, checksLine, checksLine, checksLine, checksLine)
	checks := helpers.StringToNumSlice(checksLine, ",")

	total := 0

	for i, c := range data {
		var newNewLines []string
		final := i == len(data)-1
		if c == '?' {
			for _, l := range newLines {
				newLineA := l + "."
				if final || CurrentlyValid(newLineA, checks) {
					newNewLines = append(newNewLines, newLineA)
				}

				newLineB := l + "#"
				if final || CurrentlyValid(newLineB, checks) {
					newNewLines = append(newNewLines, newLineB)
				}
			}
		} else {
			for _, l := range newLines {
				newLineA := l + string(c)
				if final || CurrentlyValid(newLineA, checks) {
					newNewLines = append(newNewLines, newLineA)
				}
			}
		}
		newLines = newNewLines
	}
	finalLines := []string{}
	for _, newLine := range newLines {
		if IsValid(newLine, checks) {
			finalLines = append(finalLines, newLine)
			total++
		}
	}
	fmt.Println(finalLines)
	ch <- total
}

func CurrentlyValid(line string, checks []int) bool {
	lineCurrentGroups := CurrentGroups(line)
	checksLength := min(len(checks), len(lineCurrentGroups))
	return slices.Equal(lineCurrentGroups, checks[:checksLength])
}

func CurrentGroups(line string) []int {
	trimmedLine := strings.TrimRight(line, "#")
	trimmedLine = strings.Trim(trimmedLine, ".")
	if len(trimmedLine) == 0 {
		return []int{}
	}
	dotRegex := regexp.MustCompile(`\.{1,}`)
	groupLines := dotRegex.Split(trimmedLine, -1)
	var groups []int
	for _, g := range groupLines {
		groups = append(groups, len(g))
	}
	return groups
}

func IsValid(line string, checks []int) bool {
	currentGroup := 0
	currentCheckIndex := 0
	for _, c := range line {
		if c == '#' {
			currentGroup++
		} else {
			if currentGroup > 0 {
				if currentCheckIndex >= len(checks) {
					return false
				}
				if currentGroup != checks[currentCheckIndex] {
					return false
				}
				currentGroup = 0
				currentCheckIndex++
			}
		}
	}
	if currentGroup > 0 {
		if currentCheckIndex >= len(checks) {
			return false
		}
		if currentGroup != checks[currentCheckIndex] {
			return false
		}
		currentCheckIndex++
	}
	return currentCheckIndex == len(checks)
}
