package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

type pos struct {
	x, y int
}

type bots struct {
	velocity pos
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")
	width := 101
	height := 103

	area := initMap(height, width)

	positionRegex := regexp.MustCompile(`^p=(\d+),(\d+)`)
	velocityRegex := regexp.MustCompile(`v=(-?\d+),(-?\d+)$`)
	for _, line := range lines {
		posData := positionRegex.FindAllSubmatch([]byte(line), 2)[0]
		x, _ := strconv.Atoi(string(posData[1]))
		y, _ := strconv.Atoi(string(posData[2]))

		velData := velocityRegex.FindAllSubmatch([]byte(line), 2)[0]
		vx, _ := strconv.Atoi(string(velData[1]))
		vy, _ := strconv.Atoi(string(velData[2]))

		area[pos{x, y}] = append(area[pos{x, y}], bots{pos{vx, vy}})
	}

	for i := 0; i < 100; i++ {
		newArea := initMap(height, width)
		for position, bots := range area {
			for _, bot := range bots {
				newX := position.x + bot.velocity.x
				newY := position.y + bot.velocity.y
				if newX < 0 {
					newX = width + newX
				}
				if newX >= width {
					newX = newX - width
				}
				if newY < 0 {
					newY = height + newY
				}
				if newY >= height {
					newY = newY - height
				}
				newArea[pos{newX, newY}] = append(newArea[pos{newX, newY}], bot)
			}
		}
		area = newArea
	}
	safetyFactor := helpers.MultSlice(getQuadrants(height, width, area))

	// Trust the magic numbers (found via visual inspection of the map until repeating frequencies emerged, lol).
	yPatternFrequency := 23
	xPatternFrequency := 48
	// Trust the magic math (for finding the intersecting frequencies).
	// 23 + 103(x) = 48 + 101(y)

	yAnswer := 0
	for ((xPatternFrequency-yPatternFrequency)+(width*yAnswer))%height != 0 {
		yAnswer++
	}
	treeSeconds := xPatternFrequency + (width * yAnswer)

	for i := 0; i < 6412; i++ {
		newArea := initMap(height, width)
		for position, bots := range area {
			for _, bot := range bots {
				newX := position.x + bot.velocity.x
				newY := position.y + bot.velocity.y
				if newX < 0 {
					newX = width + newX
				}
				if newX >= width {
					newX = newX - width
				}
				if newY < 0 {
					newY = height + newY
				}
				if newY >= height {
					newY = newY - height
				}
				newArea[pos{newX, newY}] = append(newArea[pos{newX, newY}], bot)
			}
		}
		area = newArea
	}
	printMap(height, width, area)
	fmt.Printf("Safety factor: %d\n", safetyFactor)
	fmt.Printf("Seconds to tree: %d\n", treeSeconds)
}

func initMap(height int, width int) map[pos][]bots {
	area := make(map[pos][]bots)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			area[pos{x, y}] = []bots{}
		}
	}
	return area
}

func printMap(height int, width int, area map[pos][]bots) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			count := len(area[pos{x, y}])
			if count == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("\u25A0")
			}
		}
		fmt.Println()
	}
}

func getQuadrants(height int, width int, area map[pos][]bots) []int {
	northWest := 0
	northEast := 0
	southWest := 0
	southEast := 0

	for position, bots := range area {
		if len(bots) > 0 {
			if position.x < width/2 && position.y < height/2 {
				northWest += len(bots)
			} else if position.x > width/2 && position.y < height/2 {
				northEast += len(bots)
			} else if position.x < width/2 && position.y > height/2 {
				southWest += len(bots)
			} else if position.x > width/2 && position.y > height/2 {
				southEast += len(bots)
			}
		}
	}

	return []int{northWest, northEast, southWest, southEast}
}
