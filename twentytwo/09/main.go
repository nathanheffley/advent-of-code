package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/input"
)

type coords struct {
	x int
	y int
}

var tailVisited = make(map[string]bool)

func main() {
	instructions := input.ReadInputFileToLines("input.test.txt")

	// For display purposes only
	grid := make(map[string]rune, 5*6)
	for x := 0; x < 6; x++ {
		for y := 0; y < 5; y++ {
			grid[formatCoordsForGridKey(coords{x, y})] = '.'
		}
	}
	grid[formatCoordsForGridKey(coords{x: 0, y: 4})] = 's'

	headCoords := coords{
		x: 0,
		y: 4,
	}

	tailCoords := coords{
		x: headCoords.x,
		y: headCoords.y,
	}

	tailVisited[formatCoordsForGridKey(tailCoords)] = true

	fmt.Printf("=== INITIAL STATE ===\n\n")
	printGrid(grid, headCoords, tailCoords)

	for _, instruction := range instructions {
		direction := strings.Split(instruction, " ")[0]
		distance, _ := strconv.Atoi(strings.Split(instruction, " ")[1])

		fmt.Printf("=== %s ===\n\n", instruction)

		for i := 1; i <= distance; i++ {
			if direction == "R" {
				headCoords.x++
			}
			if direction == "L" {
				headCoords.x--
			}
			if direction == "D" {
				headCoords.y++
			}
			if direction == "U" {
				headCoords.y--
			}

			tailCoords = moveTail(headCoords, tailCoords)

			printGrid(grid, headCoords, tailCoords)
		}
	}

	fmt.Print("=== Tail Visited ===\n\n")

	for key := range grid {
		if _, ok := tailVisited[key]; ok {
			grid[key] = '#'
		}
	}
	grid[formatCoordsForGridKey(coords{x: 0, y: 4})] = 's'

	printGrid(grid, coords{x: -1, y: -1}, coords{x: -1, y: -1})

	fmt.Printf("The tail visited %d spots.\n", len(tailVisited))
}

func formatCoordsForGridKey(c coords) string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func moveTail(head coords, tail coords) coords {
	xDiff := math.Abs(float64(head.x - tail.x))
	yDiff := math.Abs(float64(head.y - tail.y))

	if xDiff+yDiff == 0 {
		return tail
	}

	if xDiff < 2 && yDiff < 2 {
		return tail
	}

	if head.y > tail.y {
		tail.y++
	} else if head.y < tail.y {
		tail.y--
	}

	if head.x > tail.x {
		tail.x++
	} else if head.x < tail.x {
		tail.x--
	}

	tailVisited[formatCoordsForGridKey(tail)] = true

	return tail
}

func printGrid(grid map[string]rune, head coords, tail coords) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 6; x++ {
			key := formatCoordsForGridKey(coords{x, y})
			if key == formatCoordsForGridKey(head) {
				fmt.Print("H")
			} else if key == formatCoordsForGridKey(tail) {
				fmt.Print("T")
			} else {
				fmt.Printf("%c", grid[key])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
