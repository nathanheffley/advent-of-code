package main

import (
	"fmt"
	"math"

	"github.com/nathanheffley/advent-of-code/input"
)

type pos struct {
	x, y int
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	mapHeight := len(lines)
	mapWidth := len(lines[0])

	antennas := make(map[rune][]pos)
	for y, line := range lines {
		for x, r := range line {
			if r != '.' {
				if _, ok := antennas[r]; !ok {
					antennas[r] = make([]pos, 0)
				}
				antennas[r] = append(antennas[r], pos{x, y})
			}
		}
	}

	antinodePositions := make(map[pos]bool, 0)
	superAntinodePositions := make(map[pos]bool, 0)
	for _, positions := range antennas {
		for i, p1 := range positions {
			superAntinodePositions[p1] = true
			for _, p2 := range positions[i+1:] {
				yDistance := int(math.Abs(float64(p1.y - p2.y)))
				xDistance := int(math.Abs(float64(p1.x - p2.x)))

				var p1AntinodeXDistance int
				var p2AntinodeXDistance int
				if p1.x < p2.x {
					p1AntinodeXDistance = -xDistance
					p2AntinodeXDistance = xDistance
				} else {
					p1AntinodeXDistance = xDistance
					p2AntinodeXDistance = -xDistance
				}

				var p1AntinodeYDistance int
				var p2AntinodeYDistance int
				if p1.y < p2.y {
					p1AntinodeYDistance = -yDistance
					p2AntinodeYDistance = yDistance
				} else {
					p1AntinodeYDistance = yDistance
					p2AntinodeYDistance = -yDistance
				}

				p1AntinodeX := p1.x + p1AntinodeXDistance
				p1AntinodeY := p1.y + p1AntinodeYDistance
				if p1AntinodeX >= 0 && p1AntinodeX < mapWidth && p1AntinodeY >= 0 && p1AntinodeY < mapHeight {
					antinodePositions[pos{p1AntinodeX, p1AntinodeY}] = true
				}
				for p1AntinodeX >= 0 && p1AntinodeX < mapWidth && p1AntinodeY >= 0 && p1AntinodeY < mapHeight {
					superAntinodePositions[pos{p1AntinodeX, p1AntinodeY}] = true
					p1AntinodeX += p1AntinodeXDistance
					p1AntinodeY += p1AntinodeYDistance
				}

				p2AntinodeX := p2.x + p2AntinodeXDistance
				p2AntinodeY := p2.y + p2AntinodeYDistance
				if p2AntinodeX >= 0 && p2AntinodeX < mapWidth && p2AntinodeY >= 0 && p2AntinodeY < mapHeight {
					antinodePositions[pos{p2AntinodeX, p2AntinodeY}] = true
				}
				for p2AntinodeX >= 0 && p2AntinodeX < mapWidth && p2AntinodeY >= 0 && p2AntinodeY < mapHeight {
					superAntinodePositions[pos{p2AntinodeX, p2AntinodeY}] = true
					p2AntinodeX += p2AntinodeXDistance
					p2AntinodeY += p2AntinodeYDistance
				}
			}
		}
	}

	fmt.Printf("Antinode count: %d\n", len(antinodePositions))
	fmt.Printf("Super antinode count: %d\n", len(superAntinodePositions))
}
