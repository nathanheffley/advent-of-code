package main

import (
	"fmt"
	"math"

	"github.com/nathanheffley/advent-of-code/input"
)

type coords struct {
	x, y int
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	expandedRowIndices := make(map[int]int)
	expandedColumnIndices := make(map[int]int)
	// expansion := 2
	expansion := 1000000

	for y := 0; y < len(lines); y++ {
		isEmptyRow := true
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == '#' {
				isEmptyRow = false
				break
			}
		}
		if isEmptyRow {
			expandedRowIndices[y] = expansion
		}
	}
	for x := 0; x < len(lines[0]); x++ {
		isEmptyColumn := true
		for y := 0; y < len(lines); y++ {
			if lines[y][x] == '#' {
				isEmptyColumn = false
				break
			}
		}
		if isEmptyColumn {
			expandedColumnIndices[x] = expansion
		}
	}

	var galaxies []coords
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				galaxies = append(galaxies, coords{x, y})
			}
		}
	}

	var allPairs [][2]coords
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			allPairs = append(allPairs, [2]coords{galaxies[i], galaxies[j]})
		}
	}

	totalDistance := 0
	for _, pair := range allPairs {
		smallX := int(math.Min(float64(pair[0].x), float64(pair[1].x)))
		largeX := int(math.Max(float64(pair[0].x), float64(pair[1].x)))
		smallY := int(math.Min(float64(pair[0].y), float64(pair[1].y)))
		largeY := int(math.Max(float64(pair[0].y), float64(pair[1].y)))

		xDist := 0
		for x := smallX; x < int(largeX); x++ {
			if _, ok := expandedColumnIndices[x]; ok {
				xDist += expansion
			} else {
				xDist++
			}
		}

		yDist := 0
		for y := smallY; y < int(largeY); y++ {
			if _, ok := expandedRowIndices[y]; ok {
				yDist += expansion
			} else {
				yDist++
			}
		}

		totalDistance += xDist + yDist
	}

	fmt.Println(totalDistance)
}
