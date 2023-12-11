package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/input"
)

var layout = make(map[int][]string)

var allPipePieces = make(map[[2]int]string)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	// [y][x]
	startingCoords := [2]int{0, 0}
	for y, line := range lines {
		for x, char := range line {
			layout[y] = append(layout[y], string(char))
			if char == 'S' {
				startingCoords = [2]int{y, x}
			}
		}
	}

	firstPipeCoords, direction, char := getConnectedPipeCoords(startingCoords)
	allPipePieces[[2]int{startingCoords[0], startingCoords[1]}] = "S"
	allPipePieces[firstPipeCoords] = char
	buildPipe(firstPipeCoords, direction)

	maxSteps := (len(allPipePieces)) / 2

	fmt.Printf("Part 1 Steps: %d\n", maxSteps)

	isInside := false
	insideTiles := 0
	for y := 0; y < len(layout); y++ {
		for x := 0; x < len(layout[y]); x++ {
			char, ok := allPipePieces[[2]int{y, x}]

			if !ok {
				if isInside {
					insideTiles++
				}
			} else if char == "S" || char == "F" || char == "7" || char == "|" {
				isInside = !isInside
			}

			// Optional Visualization
			// if ok {
			// 	fmt.Print("\033[34m" + char + "\033[0m")
			// } else if isInside {
			// 	fmt.Print("\033[31mX\033[0m")
			// } else {
			// 	fmt.Print(" ")
			// }
		}
		// fmt.Println()
	}

	fmt.Printf("Part 2 Area: %d\n", insideTiles)
}

func buildPipe(coords [2]int, direction string) {
	nextCoords, direction, char := getNextPipeCoord(coords, direction)
	_, alreadyVisited := allPipePieces[[2]int{nextCoords[0], nextCoords[1]}]
	for !alreadyVisited {
		allPipePieces[[2]int{nextCoords[0], nextCoords[1]}] = char
		nextCoords, direction, char = getNextPipeCoord(nextCoords, direction)
		_, alreadyVisited = allPipePieces[[2]int{nextCoords[0], nextCoords[1]}]
	}
}

func getNextPipeCoord(coords [2]int, incomingDirection string) ([2]int, string, string) {
	piece := layout[coords[0]][coords[1]]
	switch incomingDirection {
	case "North":
		switch piece {
		case "|":
			return [2]int{coords[0] + 1, coords[1]}, "North", layout[coords[0]+1][coords[1]]
		case "L":
			return [2]int{coords[0], coords[1] + 1}, "West", layout[coords[0]][coords[1]+1]
		case "J":
			return [2]int{coords[0], coords[1] - 1}, "East", layout[coords[0]][coords[1]-1]
		default:
			return [2]int{}, "", ""
		}
	case "South":
		switch piece {
		case "|":
			return [2]int{coords[0] - 1, coords[1]}, "South", layout[coords[0]-1][coords[1]]
		case "F":
			return [2]int{coords[0], coords[1] + 1}, "West", layout[coords[0]][coords[1]+1]
		case "7":
			return [2]int{coords[0], coords[1] - 1}, "East", layout[coords[0]][coords[1]-1]
		default:
			return [2]int{}, "", ""
		}
	case "East":
		switch piece {
		case "-":
			return [2]int{coords[0], coords[1] - 1}, "East", layout[coords[0]][coords[1]-1]
		case "F":
			return [2]int{coords[0] + 1, coords[1]}, "North", layout[coords[0]+1][coords[1]]
		case "L":
			return [2]int{coords[0] - 1, coords[1]}, "South", layout[coords[0]-1][coords[1]]
		default:
			return [2]int{}, "", ""
		}
	case "West":
		switch piece {
		case "-":
			return [2]int{coords[0], coords[1] + 1}, "West", layout[coords[0]][coords[1]+1]
		case "7":
			return [2]int{coords[0] + 1, coords[1]}, "North", layout[coords[0]+1][coords[1]]
		case "J":
			return [2]int{coords[0] - 1, coords[1]}, "South", layout[coords[0]-1][coords[1]]
		default:
			return [2]int{}, "", ""
		}
	default:
		return [2]int{}, "", ""
	}
}

// coords = [y, x], returns coord and direction coming from start
func getConnectedPipeCoords(coords [2]int) ([2]int, string, string) {
	// Up
	if coords[0] > 0 {
		piece := layout[coords[0]-1][coords[1]]
		if piece == "|" || piece == "F" || piece == "7" {
			return [2]int{coords[0] - 1, coords[1]}, "South", piece
		}
	}

	// Down
	if coords[0] < len(layout)-1 {
		piece := layout[coords[0]+1][coords[1]]
		if piece == "|" || piece == "L" || piece == "J" {
			return [2]int{coords[0] + 1, coords[1]}, "North", piece
		}
	}

	// Left
	if coords[1] > 0 {
		piece := layout[coords[0]][coords[1]-1]
		if piece == "-" || piece == "L" || piece == "F" {
			return [2]int{coords[0], coords[1] - 1}, "East", piece
		}
	}

	// Right
	if coords[1] < len(layout[coords[0]])-1 {
		piece := layout[coords[0]][coords[1]+1]
		if piece == "-" || piece == "J" || piece == "7" {
			return [2]int{coords[0], coords[1] + 1}, "West", piece
		}
	}

	return [2]int{}, "", ""
}
