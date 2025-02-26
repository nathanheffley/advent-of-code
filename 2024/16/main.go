package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/input"
)

type point struct {
	x, y int
}

type space struct {
	wall     bool
	minScore int
}

func main() {
	lines := input.ReadInputFileToLines("input.test.txt")

	start := point{0, 0}
	end := point{0, 0}
	maze := make(map[point]space)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				maze[point{x, y}] = space{true, 0}
			} else {
				maze[point{x, y}] = space{false, 0}
				if char == 'S' {
					start = point{x, y}
				} else if char == 'E' {
					end = point{x, y}
				}
			}
		}
	}

	fmt.Println(start, end)
}
