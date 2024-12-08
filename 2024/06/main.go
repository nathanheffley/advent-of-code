package main

import (
	"fmt"
	"time"

	"github.com/nathanheffley/advent-of-code/input"
)

type point struct {
	x, y                     int
	north, east, south, west bool
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	start := time.Now()

	var guardX, guardY int
	guardMap := make([][]byte, len(lines))
	for y, line := range lines {
		guardMap[y] = make([]byte, len(line))
		for x := 0; x < len(guardMap[y]); x++ {
			if line[x] == '^' {
				guardX = x
				guardY = y
				guardMap[y][x] = '.'
				continue
			}
			guardMap[y][x] = line[x]
		}
	}

	// Part one
	_, visitedLocations := walkMap(guardMap, guardX, guardY)
	fmt.Printf("Visited locations: %d\n", len(visitedLocations))

	// Part two
	loopOptions := 0
	for _, p := range visitedLocations {
		guardMap[p.y][p.x] = '#'
		isLooping, _ := walkMap(guardMap, guardX, guardY)
		if isLooping {
			loopOptions++
		}
		guardMap[p.y][p.x] = '.'
	}
	fmt.Printf("Loop options: %d\n", loopOptions)

	elapsed := time.Since(start)
	fmt.Printf("(took %s)\n", elapsed)
}

func walkMap(guardMap [][]byte, guardX int, guardY int) (bool, map[string]point) {
	guardDirection := 'N'

	visitedLocations := make(map[string]point)

	isLooping := false
guardLoop:
	for (guardX >= 0 && guardX < len(guardMap[0])) && (guardY >= 0 && guardY < len(guardMap)) {
		key := fmt.Sprintf("%d,%d", guardX, guardY)
		if _, ok := visitedLocations[key]; !ok {
			visitedLocations[key] = point{
				x:     guardX,
				y:     guardY,
				north: false,
				east:  false,
				south: false,
				west:  false,
			}
		}

		switch guardDirection {
		case 'N':
			if visitedLocations[key].north {
				isLooping = true
				break guardLoop
			}
			visitedLocations[key] = point{
				x:     guardX,
				y:     guardY,
				north: true,
				east:  visitedLocations[key].east,
				south: visitedLocations[key].south,
				west:  visitedLocations[key].west,
			}
			if guardY > 0 && guardMap[guardY-1][guardX] == '#' {
				guardDirection = 'E'
			} else {
				guardY--
			}
		case 'E':
			if visitedLocations[key].east {
				isLooping = true
				break guardLoop
			}
			visitedLocations[key] = point{
				x:     guardX,
				y:     guardY,
				north: visitedLocations[key].north,
				east:  true,
				south: visitedLocations[key].south,
				west:  visitedLocations[key].west,
			}
			if guardX < len(guardMap[0])-1 && guardMap[guardY][guardX+1] == '#' {
				guardDirection = 'S'
			} else {
				guardX++
			}
		case 'S':
			if visitedLocations[key].south {
				isLooping = true
				break guardLoop
			}
			visitedLocations[key] = point{
				x:     guardX,
				y:     guardY,
				north: visitedLocations[key].north,
				east:  visitedLocations[key].east,
				south: true,
				west:  visitedLocations[key].west,
			}
			if guardY < len(guardMap)-1 && guardMap[guardY+1][guardX] == '#' {
				guardDirection = 'W'
			} else {
				guardY++
			}
		case 'W':
			if visitedLocations[key].west {
				isLooping = true
				break guardLoop
			}
			visitedLocations[key] = point{
				x:     guardX,
				y:     guardY,
				north: visitedLocations[key].north,
				east:  visitedLocations[key].east,
				south: visitedLocations[key].south,
				west:  true,
			}
			if guardX > 0 && guardMap[guardY][guardX-1] == '#' {
				guardDirection = 'N'
			} else {
				guardX--
			}
		}
	}

	return isLooping, visitedLocations
}
