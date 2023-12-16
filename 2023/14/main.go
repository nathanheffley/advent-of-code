package main

import (
	"fmt"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

var fullCycleCache = make(map[string][]string)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	part1Lines := cycleNorth(lines)
	fmt.Printf("Part 1 Total: %d\n", scoreLines(part1Lines))

	part2Lines := append([]string{}, lines...)
	for i := 0; i < 1000; i++ {
		cacheKey := strings.Join(part2Lines, "")
		if val, ok := fullCycleCache[cacheKey]; ok {
			part2Lines = val
			continue
		}
		part2Lines = cycleNorth(part2Lines)
		part2Lines = cycleWest(part2Lines)
		part2Lines = cycleSouth(part2Lines)
		part2Lines = cycleEast(part2Lines)
		fullCycleCache[cacheKey] = part2Lines
	}

	fmt.Printf("Part 2 Total: %d\n", scoreLines(part2Lines))
}

func scoreLines(lines []string) int {
	total := 0
	for y := 0; y < len(lines); y++ {
		for _, c := range lines[y] {
			if c == 'O' {
				total += len(lines) - y
			}
		}
	}
	return total
}

var shiftLineCache = make(map[string]string)

func shiftLine(line string) string {
	if val, ok := shiftLineCache[line]; ok {
		return val
	}
	array := []rune(line)
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] == '.' && array[j+1] == 'O' {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	shiftLineCache[line] = string(array)
	return string(array)
}

var cycleNorthCache = make(map[string][]string)

func cycleNorth(lines []string) []string {
	if result, ok := cycleNorthCache[strings.Join(lines, "")]; ok {
		return result
	}
	linesToShift := []string{}
	for x := 0; x < len(lines[0]); x++ {
		line := ""
		for y := 0; y < len(lines); y++ {
			line += string(lines[y][x])
		}
		linesToShift = append(linesToShift, shiftLine(line))
	}
	// Rotate the lines back to the northward direction
	newLines := []string{}
	for x := 0; x < len(linesToShift[0]); x++ {
		line := ""
		for y := 0; y < len(linesToShift); y++ {
			line += string(linesToShift[y][x])
		}
		newLines = append(newLines, line)
	}
	cycleNorthCache[strings.Join(lines, "")] = newLines
	return newLines
}

var cycleWestCache = make(map[string][]string)

func cycleWest(lines []string) []string {
	if result, ok := cycleWestCache[strings.Join(lines, "")]; ok {
		return result
	}
	newLines := []string{}
	for _, line := range lines {
		newLines = append(newLines, shiftLine(line))
	}
	cycleWestCache[strings.Join(lines, "")] = newLines
	return newLines
}

var cycleSouthCache = make(map[string][]string)

func cycleSouth(lines []string) []string {
	if result, ok := cycleSouthCache[strings.Join(lines, "")]; ok {
		return result
	}
	linesToShift := []string{}
	for x := 0; x < len(lines[0]); x++ {
		line := ""
		for y := 0; y < len(lines); y++ {
			line += string(lines[y][x])
		}
		line = helpers.Reverse(line)
		line = shiftLine(line)
		line = helpers.Reverse(line)
		linesToShift = append(linesToShift, line)
	}
	// Rotate the lines back to the southward direction
	newLines := []string{}
	for x := 0; x < len(linesToShift[0]); x++ {
		line := ""
		for y := 0; y < len(linesToShift); y++ {
			line += string(linesToShift[y][x])
		}
		newLines = append(newLines, line)
	}
	cycleSouthCache[strings.Join(lines, "")] = newLines
	return newLines
}

var cycleEastCache = make(map[string][]string)

func cycleEast(lines []string) []string {
	if result, ok := cycleEastCache[strings.Join(lines, "")]; ok {
		return result
	}
	newLines := []string{}
	for _, line := range lines {
		line = helpers.Reverse(line)
		line = shiftLine(line)
		line = helpers.Reverse(line)
		newLines = append(newLines, line)
	}
	cycleEastCache[strings.Join(lines, "")] = newLines
	return newLines
}
