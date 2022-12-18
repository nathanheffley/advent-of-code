package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/input"
)

type coordinate struct {
	x int
	y int
}

func (c *coordinate) key() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func main() {
	data := input.ReadInputFileToLines("input.test.txt")

	solids := make(map[string]string)
	solidsWithFloor := make(map[string]string)

	smallestX := math.MaxInt
	biggestX := 0
	biggestY := 0

	for _, datum := range data {
		lines := strings.Split(datum, " -> ")
		coords := []coordinate{}
		for _, line := range lines {
			x, _ := strconv.Atoi(strings.Split(line, ",")[0])
			y, _ := strconv.Atoi(strings.Split(line, ",")[1])

			if x < smallestX {
				smallestX = x
			}
			if x > biggestX {
				biggestX = x
			}
			if y > biggestY {
				biggestY = y
			}

			coords = append(coords, coordinate{x, y})
		}
		for i := 0; i < len(coords)-1; i++ {
			start := coords[i]
			end := coords[i+1]

			for x := start.x; x <= end.x; x++ {
				for y := start.y; y <= end.y; y++ {
					coord := coordinate{x, y}
					solids[coord.key()] = "#"
					solidsWithFloor[coord.key()] = "#"
				}
			}
			for x := end.x; x <= start.x; x++ {
				for y := end.y; y <= start.y; y++ {
					coord := coordinate{x, y}
					solids[coord.key()] = "#"
					solidsWithFloor[coord.key()] = "#"
				}
			}
		}
	}

	for x := smallestX - 1000; x < biggestX+1000; x++ {
		coord := coordinate{
			x: x,
			y: biggestY + 2,
		}
		solidsWithFloor[coord.key()] = "#"
	}

	sandCount := 0
	for {
		fellIntoAbyss := dropSand(solids, biggestY)
		if fellIntoAbyss {
			break
		}
		sandCount++
	}
	fmt.Printf("Part 1: %d\n", sandCount)
	drawMap(solids)
	fmt.Println()

	sandCount = 0
	spawnCoord := coordinate{
		x: 500,
		y: 0,
	}
	for {
		if _, solid := solidsWithFloor[spawnCoord.key()]; solid {
			break
		}
		fellIntoAbyss := dropSand(solidsWithFloor, biggestY+2)
		if fellIntoAbyss {
			panic("Whoopsie")
		}
		sandCount++
	}
	fmt.Printf("Part 2: %d\n", sandCount)
	drawMap(solidsWithFloor)
}

func dropSand(solids map[string]string, abyssY int) bool {
	sandCoord := coordinate{
		x: 500,
		y: 0,
	}
	for {
		checkCoord := coordinate{
			x: sandCoord.x,
			y: sandCoord.y + 1,
		}

		if _, solid := solids[checkCoord.key()]; solid {
			leftCheckCoord := coordinate{
				x: checkCoord.x - 1,
				y: checkCoord.y,
			}
			if _, solidLeft := solids[leftCheckCoord.key()]; !solidLeft {
				sandCoord = coordinate{
					x: leftCheckCoord.x,
					y: leftCheckCoord.y,
				}
				continue
			}

			rightCheckCoord := coordinate{
				x: checkCoord.x + 1,
				y: checkCoord.y,
			}
			if _, solidRight := solids[rightCheckCoord.key()]; !solidRight {
				sandCoord = coordinate{
					x: rightCheckCoord.x,
					y: rightCheckCoord.y,
				}
				continue
			}

			// Sand is stuck
			solids[sandCoord.key()] = "o"
			break
		}

		sandCoord = coordinate{
			x: checkCoord.x,
			y: checkCoord.y,
		}

		if sandCoord.y > abyssY {
			return true
		}
	}

	return false
}

func drawMap(solids map[string]string) {
	for y := 0; y <= 11; y++ {
		for x := 488; x <= 512; x++ {
			coord := coordinate{x, y}
			if s, ok := solids[coord.key()]; ok {
				fmt.Print(s)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
