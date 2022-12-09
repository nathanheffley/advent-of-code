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

type knot struct {
	char    string
	coords  coords
	visited map[string]bool
}

func main() {
	instructions := input.ReadInputFileToLines("input.txt")

	// For test display purposes only
	grid := make(map[string]string, 5*6)
	for x := 0; x < 6; x++ {
		for y := 0; y < 5; y++ {
			grid[formatCoordsForGridKey(coords{x, y})] = "."
		}
	}
	grid[formatCoordsForGridKey(coords{x: 0, y: 4})] = "s"

	head := knot{
		char: "H",
		coords: coords{
			x: 0,
			y: 4,
		},
		visited: map[string]bool{
			"0,4": true,
		},
	}

	tailOne := knot{
		char: "1",
		coords: coords{
			x: 0,
			y: 4,
		},
		visited: map[string]bool{
			"0,4": true,
		},
	}

	tailTwo := knot{
		char: "2",
		coords: coords{
			x: 0,
			y: 4,
		},
		visited: map[string]bool{
			"0,4": true,
		},
	}

	tailThree := knot{
		char: "3",
		coords: coords{
			x: 0,
			y: 4,
		},
		visited: map[string]bool{
			"0,4": true,
		},
	}

	tailFour := knot{
		char: "4",
		coords: coords{
			x: 0,
			y: 4,
		},
		visited: map[string]bool{
			"0,4": true,
		},
	}

	tailFive := knot{
		char: "5",
		coords: coords{
			x: 0,
			y: 4,
		},
		visited: map[string]bool{
			"0,4": true,
		},
	}

	tailSix := knot{
		char: "6",
		coords: coords{
			x: 0,
			y: 4,
		},
		visited: map[string]bool{
			"0,4": true,
		},
	}

	tailSeven := knot{
		char: "7",
		coords: coords{
			x: 0,
			y: 4,
		},
		visited: map[string]bool{
			"0,4": true,
		},
	}

	tailEight := knot{
		char: "8",
		coords: coords{
			x: 0,
			y: 4,
		},
		visited: map[string]bool{
			"0,4": true,
		},
	}

	tailNine := knot{
		char: "9",
		coords: coords{
			x: 0,
			y: 4,
		},
		visited: map[string]bool{
			"0,4": true,
		},
	}

	// fmt.Printf("=== INITIAL STATE ===\n\n")
	// printGrid(grid, head, tailOne, tailTwo)

	for _, instruction := range instructions {
		direction := strings.Split(instruction, " ")[0]
		distance, _ := strconv.Atoi(strings.Split(instruction, " ")[1])

		// fmt.Printf("=== %s ===\n\n", instruction)

		for i := 1; i <= distance; i++ {
			if direction == "R" {
				head.coords.x++
			}
			if direction == "L" {
				head.coords.x--
			}
			if direction == "D" {
				head.coords.y++
			}
			if direction == "U" {
				head.coords.y--
			}

			tailOne = moveTail(head, tailOne)

			tailTwo = moveTail(tailOne, tailTwo)

			tailThree = moveTail(tailTwo, tailThree)

			tailFour = moveTail(tailThree, tailFour)

			tailFive = moveTail(tailFour, tailFive)

			tailSix = moveTail(tailFive, tailSix)

			tailSeven = moveTail(tailSix, tailSeven)

			tailEight = moveTail(tailSeven, tailEight)

			tailNine = moveTail(tailEight, tailNine)

			// printGrid(grid, head, tailOne, tailTwo)
		}
	}

	fmt.Printf("The first tail visited %d spots.\n", len(tailOne.visited))

	fmt.Printf("The ninth tail visited %d spots.\n", len(tailNine.visited))
}

func formatCoordsForGridKey(c coords) string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func moveTail(head knot, tail knot) knot {
	xDiff := math.Abs(float64(head.coords.x - tail.coords.x))
	yDiff := math.Abs(float64(head.coords.y - tail.coords.y))

	if xDiff+yDiff == 0 {
		return tail
	}

	if xDiff < 2 && yDiff < 2 {
		return tail
	}

	if head.coords.y > tail.coords.y {
		tail.coords.y++
	} else if head.coords.y < tail.coords.y {
		tail.coords.y--
	}

	if head.coords.x > tail.coords.x {
		tail.coords.x++
	} else if head.coords.x < tail.coords.x {
		tail.coords.x--
	}

	tail.visited[formatCoordsForGridKey(tail.coords)] = true

	return tail
}

func printGrid(grid map[string]string, head knot, tail knot, tailTwo knot) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 6; x++ {
			key := formatCoordsForGridKey(coords{x, y})
			if key == formatCoordsForGridKey(head.coords) {
				fmt.Print(head.char)
			} else if key == formatCoordsForGridKey(tail.coords) {
				fmt.Print(tail.char)
			} else if key == formatCoordsForGridKey(tailTwo.coords) {
				fmt.Print(tailTwo.char)
			} else {
				fmt.Printf("%s", grid[key])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
