package main

import (
	"fmt"
	"math"

	"github.com/nathanheffley/advent-of-code/input"
)

type coordinate struct {
	x       int
	y       int
	checked map[string]rune
}

func (c *coordinate) key() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

var topoMap map[string]int

var mapWidth int
var mapHeight int
var goalCoord coordinate

var minTravel = math.MaxInt
var minTravelNode coordinate

var minTravelToNode = make(map[string]int)

var nextToCheck []coordinate

func main() {
	data := input.ReadInputFileToLines("input.txt")

	mapWidth = len(data[0])
	mapHeight = len(data)

	topoMap = make(map[string]int)

	possibleStartingPoints := []coordinate{}

	var startCoord coordinate

	for y, line := range data {
		for x, point := range line {
			coord := coordinate{
				x:       x,
				y:       y,
				checked: make(map[string]rune),
			}

			elevation := int(point)
			if point == 'a' {
				possibleStartingPoints = append(possibleStartingPoints, coord)
			}
			if point == 'S' {
				startCoord = coord
				elevation = int('a')
				possibleStartingPoints = append(possibleStartingPoints, coord)
			}
			if point == 'E' {
				goalCoord = coord
				elevation = int('z')
			}

			topoMap[coord.key()] = elevation
		}
	}

	nextToCheck = append(nextToCheck, startCoord)
	checkPoints()
	fmt.Printf("Part 1: %d\n", minTravel)
	printMap(minTravelNode)

	part2MinTravel := math.MaxInt
	for _, c := range possibleStartingPoints {
		minTravelToNode = make(map[string]int)
		nextToCheck = []coordinate{c}
		checkPoints()
		if minTravel < part2MinTravel {
			part2MinTravel = minTravel
		}
	}
	fmt.Printf("\nPart 2: %d\n", part2MinTravel)
	printMap(minTravelNode)
}

func checkPoints() {
	newNextToCheck := []coordinate{}

	for _, p := range nextToCheck {
		newNextToCheck = append(newNextToCheck, checkPoint(p)...)
	}

	if len(newNextToCheck) > 0 {
		nextToCheck = newNextToCheck
		checkPoints()
	}
}

func checkPoint(checkCoord coordinate) []coordinate {
	if checkCoord.x < 0 || checkCoord.x >= mapWidth {
		return []coordinate{}
	}
	if checkCoord.y < 0 || checkCoord.y >= mapHeight {
		return []coordinate{}
	}
	if _, checked := checkCoord.checked[checkCoord.key()]; checked {
		return []coordinate{}
	}

	if shortest, ok := minTravelToNode[checkCoord.key()]; ok {
		if len(checkCoord.checked) >= shortest {
			return []coordinate{}
		} else {
			minTravelToNode[checkCoord.key()] = len(checkCoord.checked)
		}
	} else {
		minTravelToNode[checkCoord.key()] = len(checkCoord.checked)
	}

	if checkCoord.key() == goalCoord.key() {
		if len(checkCoord.checked) < minTravel {
			minTravel = len(checkCoord.checked)
			minTravelNode = checkCoord
		}
		return []coordinate{}
	}

	newPointsToCheck := []coordinate{}

	up := coordinate{
		x:       checkCoord.x,
		y:       checkCoord.y - 1,
		checked: nil,
	}
	if topoMap[up.key()] <= (topoMap[checkCoord.key()] + 1) {
		newCheckedMap := make(map[string]rune, len(checkCoord.checked))
		for k, v := range checkCoord.checked {
			newCheckedMap[k] = v
		}
		newCheckedMap[checkCoord.key()] = '^'
		up.checked = newCheckedMap
		newPointsToCheck = append(newPointsToCheck, up)
	}

	down := coordinate{
		x:       checkCoord.x,
		y:       checkCoord.y + 1,
		checked: nil,
	}
	if topoMap[down.key()] <= (topoMap[checkCoord.key()] + 1) {
		newCheckedMap := make(map[string]rune, len(checkCoord.checked))
		for k, v := range checkCoord.checked {
			newCheckedMap[k] = v
		}
		newCheckedMap[checkCoord.key()] = 'V'
		down.checked = newCheckedMap
		newPointsToCheck = append(newPointsToCheck, down)
	}

	left := coordinate{
		x:       checkCoord.x - 1,
		y:       checkCoord.y,
		checked: nil,
	}
	if topoMap[left.key()] <= (topoMap[checkCoord.key()] + 1) {
		newCheckedMap := make(map[string]rune, len(checkCoord.checked))
		for k, v := range checkCoord.checked {
			newCheckedMap[k] = v
		}
		newCheckedMap[checkCoord.key()] = '<'
		left.checked = newCheckedMap
		newPointsToCheck = append(newPointsToCheck, left)
	}

	right := coordinate{
		x:       checkCoord.x + 1,
		y:       checkCoord.y,
		checked: nil,
	}
	if topoMap[right.key()] <= (topoMap[checkCoord.key()] + 1) {
		newCheckedMap := make(map[string]rune, len(checkCoord.checked))
		for k, v := range checkCoord.checked {
			newCheckedMap[k] = v
		}
		newCheckedMap[checkCoord.key()] = '>'
		right.checked = newCheckedMap
		newPointsToCheck = append(newPointsToCheck, right)
	}

	return newPointsToCheck
}

func printMap(coord coordinate) {
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			coord := coordinate{
				x:       x,
				y:       y,
				checked: coord.checked,
			}
			if coord.key() == goalCoord.key() {
				fmt.Print("E")
			} else if r, ok := coord.checked[coord.key()]; ok {
				fmt.Printf("%c", r)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
