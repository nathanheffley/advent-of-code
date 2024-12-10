package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/input"
)

type position struct {
	x, y int
}

type location struct {
	position position
	height   int
	// rating   int
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	maze := make(map[position]location)
	startingLocations := make(map[position]location, 0)
	for y, line := range lines {
		for x, char := range line {
			pos := position{x, y}
			loc := location{pos, int(char - '0')}
			// loc := location{pos, int(char - '0'), -1}
			maze[pos] = loc
			if loc.height == 0 {
				// loc.rating = 1
				startingLocations[pos] = loc
			}
		}
	}

	totalEnds := 0
	totalPaths := 0
	for _, loc := range startingLocations {
		ends := walkTrail(loc, maze)
		totalPaths += len(ends)
		uniqueEnds := make(map[position]bool)
		for _, end := range ends {
			uniqueEnds[end.position] = true
		}
		totalEnds += len(uniqueEnds)
	}

	fmt.Println("Total ends:", totalEnds)
	fmt.Println("Total paths:", totalPaths)
}

func walkTrail(start location, maze map[position]location) []location {
	startingLocations := []location{start}
	ends := stepUp(startingLocations, maze)
	for i := 1; i < 9; i++ {
		ends = stepUp(ends, maze)
	}
	return ends
}

func stepUp(locs []location, maze map[position]location) []location {
	newLocations := make([]location, 0)
	for _, loc := range locs {
		// North
		northPosition := position{loc.position.x, loc.position.y - 1}
		if northLocation, ok := maze[northPosition]; ok {
			if northLocation.height == loc.height+1 {
				newLocations = append(newLocations, northLocation)
			}
		}

		// East
		eastPosition := position{loc.position.x + 1, loc.position.y}
		if eastLocation, ok := maze[eastPosition]; ok {
			if eastLocation.height == loc.height+1 {
				newLocations = append(newLocations, eastLocation)
			}
		}

		// South
		southPosition := position{loc.position.x, loc.position.y + 1}
		if southLocation, ok := maze[southPosition]; ok {
			if southLocation.height == loc.height+1 {
				newLocations = append(newLocations, southLocation)
			}
		}

		// West
		westPosition := position{loc.position.x - 1, loc.position.y}
		if westLocation, ok := maze[westPosition]; ok {
			if westLocation.height == loc.height+1 {
				newLocations = append(newLocations, westLocation)
			}
		}
	}
	return newLocations
}
