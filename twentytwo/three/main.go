package main

import (
	"fmt"
	"strings"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	priority := 0

	for _, line := range lines {
		side1 := line[:len(line)/2]
		side2 := line[len(line)/2:]

		var matches rune
	Loop:
		for _, a := range side1 {
			for _, b := range side2 {
				if a == b {
					matches = a
					break Loop
				}
			}
		}
		if matches > 90 {
			priority += int(matches) - 96
		} else {
			priority += int(matches) - 38
		}
	}

	fmt.Printf("Misplaced Items Priority: %d\n", priority)

	// --------

	badgePriority := 0

	var groups [][]string
	for i := 0; i < len(lines); i += 3 {
		groups = append(groups, lines[i:i+3])
	}
	for _, group := range groups {
		first := group[0]
		second := group[1]
		third := group[2]
		for _, c := range first {
			if strings.Contains(second, string(c)) && strings.Contains(third, string(c)) {
				if c > 90 {
					badgePriority += int(c) - 96
				} else {
					badgePriority += int(c) - 38
				}
				break
			}
		}
	}

	fmt.Printf("Badge Priority: %d\n", badgePriority)
}
