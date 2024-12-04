package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	grid := gridFromLines(lines)

	xmases := 0
	mass := 0
	keepMap := make(map[int]map[int]bool)
	masKeepMap := make(map[int]map[int]bool)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'X' {
				xmasesFound, keepMapFound := findXMASAroundPoint(grid, x, y)
				xmases += xmasesFound
				for _, point := range keepMapFound {
					if keepMap[point[0]] == nil {
						keepMap[point[0]] = make(map[int]bool)
					}
					keepMap[point[0]][point[1]] = true
				}
			}

			if grid[y][x] == 'A' {
				massFound, masKeepMapFound := findMASsAroundPoint(grid, x, y)
				mass += massFound
				for _, point := range masKeepMapFound {
					if masKeepMap[point[0]] == nil {
						masKeepMap[point[0]] = make(map[int]bool)
					}
					masKeepMap[point[0]][point[1]] = true
				}
			}
		}
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if masKeepMap[y][x] {
				fmt.Print(string(grid[y][x]))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Printf("XMAS count: %d\n", xmases)
	fmt.Printf("MAS count: %d\n", mass)
}

func gridFromLines(lines []string) [][]rune {
	grid := make([][]rune, len(lines))

	for y, line := range lines {
		grid[y] = make([]rune, len(line))
		for x, char := range line {
			grid[y][x] = char
		}
	}

	return grid
}

func findXMASAroundPoint(grid [][]rune, x int, y int) (int, [][]int) {
	canCheckUp := y >= 3
	canCheckDown := y < len(grid)-3
	canCheckLeft := x >= 3
	canCheckRight := x < len(grid[y])-3

	total := 0
	keepMap := make([][]int, 0)

	if canCheckUp && grid[y-1][x] == 'M' && grid[y-2][x] == 'A' && grid[y-3][x] == 'S' {
		total += 1
		keepMap = append(keepMap, []int{y, x})
		keepMap = append(keepMap, []int{y - 1, x})
		keepMap = append(keepMap, []int{y - 2, x})
		keepMap = append(keepMap, []int{y - 3, x})
	}

	if canCheckDown && grid[y+1][x] == 'M' && grid[y+2][x] == 'A' && grid[y+3][x] == 'S' {
		total += 1
		keepMap = append(keepMap, []int{y, x})
		keepMap = append(keepMap, []int{y + 1, x})
		keepMap = append(keepMap, []int{y + 2, x})
		keepMap = append(keepMap, []int{y + 3, x})
	}

	if canCheckLeft && grid[y][x-1] == 'M' && grid[y][x-2] == 'A' && grid[y][x-3] == 'S' {
		total += 1
		keepMap = append(keepMap, []int{y, x})
		keepMap = append(keepMap, []int{y, x - 1})
		keepMap = append(keepMap, []int{y, x - 2})
		keepMap = append(keepMap, []int{y, x - 3})
	}

	if canCheckRight && grid[y][x+1] == 'M' && grid[y][x+2] == 'A' && grid[y][x+3] == 'S' {
		total += 1
		keepMap = append(keepMap, []int{y, x})
		keepMap = append(keepMap, []int{y, x + 1})
		keepMap = append(keepMap, []int{y, x + 2})
		keepMap = append(keepMap, []int{y, x + 3})
	}

	if canCheckUp && canCheckLeft && grid[y-1][x-1] == 'M' && grid[y-2][x-2] == 'A' && grid[y-3][x-3] == 'S' {
		total += 1
		keepMap = append(keepMap, []int{y, x})
		keepMap = append(keepMap, []int{y - 1, x - 1})
		keepMap = append(keepMap, []int{y - 2, x - 2})
		keepMap = append(keepMap, []int{y - 3, x - 3})
	}

	if canCheckUp && canCheckRight && grid[y-1][x+1] == 'M' && grid[y-2][x+2] == 'A' && grid[y-3][x+3] == 'S' {
		total += 1
		keepMap = append(keepMap, []int{y, x})
		keepMap = append(keepMap, []int{y - 1, x + 1})
		keepMap = append(keepMap, []int{y - 2, x + 2})
		keepMap = append(keepMap, []int{y - 3, x + 3})
	}

	if canCheckDown && canCheckLeft && grid[y+1][x-1] == 'M' && grid[y+2][x-2] == 'A' && grid[y+3][x-3] == 'S' {
		total += 1
		keepMap = append(keepMap, []int{y, x})
		keepMap = append(keepMap, []int{y + 1, x - 1})
		keepMap = append(keepMap, []int{y + 2, x - 2})
		keepMap = append(keepMap, []int{y + 3, x - 3})
	}

	if canCheckDown && canCheckRight && grid[y+1][x+1] == 'M' && grid[y+2][x+2] == 'A' && grid[y+3][x+3] == 'S' {
		total += 1
		keepMap = append(keepMap, []int{y, x})
		keepMap = append(keepMap, []int{y + 1, x + 1})
		keepMap = append(keepMap, []int{y + 2, x + 2})
		keepMap = append(keepMap, []int{y + 3, x + 3})
	}

	return total, keepMap
}

func findMASsAroundPoint(grid [][]rune, x int, y int) (int, [][]int) {
	canCheckUp := y >= 1
	canCheckDown := y < len(grid)-1
	canCheckLeft := x >= 1
	canCheckRight := x < len(grid[y])-1

	if !canCheckUp || !canCheckDown || !canCheckLeft || !canCheckRight {
		return 0, [][]int{}
	}

	total := 0
	keepMap := make([][]int, 0)

	if (grid[y+1][x-1] == 'M' && grid[y-1][x+1] == 'S') || (grid[y+1][x-1] == 'S' && grid[y-1][x+1] == 'M') {
		if (grid[y+1][x+1] == 'M' && grid[y-1][x-1] == 'S') || (grid[y+1][x+1] == 'S' && grid[y-1][x-1] == 'M') {
			total += 1
			keepMap = append(keepMap, []int{y, x})
			keepMap = append(keepMap, []int{y - 1, x + 1})
			keepMap = append(keepMap, []int{y + 1, x - 1})
			keepMap = append(keepMap, []int{y + 1, x + 1})
			keepMap = append(keepMap, []int{y - 1, x - 1})
		}
	}

	return total, keepMap
}
