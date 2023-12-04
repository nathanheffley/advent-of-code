package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	rows := input.ReadInputFileToLines("input.txt")

	width := len(rows[0])

	total := 0

	// key will be "x,y", value is array of adjacent numbers
	gearNums := make(map[string][]int)

	for y, row := range rows {
		numberRegex := regexp.MustCompile(`\d+`)
		numberIndices := numberRegex.FindAllStringIndex(row, -1)
		for _, indices := range numberIndices {
			num, err := strconv.Atoi(row[indices[0]:indices[1]])
			helpers.Check(err)

			if indices[0] > 0 {
				if rows[y][indices[0]-1] == '*' {
					gearKey := fmt.Sprintf("%d,%d", indices[0]-1, y)
					gearNums[gearKey] = append(gearNums[gearKey], num)
				}
				if rows[y][indices[0]-1] != '.' {
					total += num
					continue
				}
			}

			if indices[1] < width {
				if rows[y][indices[1]] == '*' {
					gearKey := fmt.Sprintf("%d,%d", indices[1], y)
					gearNums[gearKey] = append(gearNums[gearKey], num)
				}
				if rows[y][indices[1]] != '.' {
					total += num
					continue
				}
			}

			foundSymbol := false
			if y > 0 {
				// check above
				for x := indices[0] - 1; x < indices[1]+1; x++ {
					if x < 0 || x >= width {
						continue
					}
					if rows[y-1][x] == '*' {
						gearKey := fmt.Sprintf("%d,%d", x, y-1)
						gearNums[gearKey] = append(gearNums[gearKey], num)
					}
					if rows[y-1][x] != '.' {
						foundSymbol = true
						break
					}
				}
			}
			if y < len(rows)-1 {
				// check below
				for x := indices[0] - 1; x < indices[1]+1; x++ {
					if x < 0 || x >= width {
						continue
					}
					if rows[y+1][x] == '*' {
						gearKey := fmt.Sprintf("%d,%d", x, y+1)
						gearNums[gearKey] = append(gearNums[gearKey], num)
					}
					if rows[y+1][x] != '.' {
						foundSymbol = true
						break
					}
				}
			}
			if foundSymbol {
				total += num
				continue
			}
		}
	}

	gearTotal := 0
	for _, nums := range gearNums {
		if len(nums) == 2 {
			gearTotal += nums[0] * nums[1]
		}
	}

	fmt.Printf("Part 1: %d\n", total)
	fmt.Printf("Part 2: %d\n", gearTotal)
}
