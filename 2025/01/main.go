package main

import (
	"fmt"
	"strconv"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	num := 50
	pointAtZeroCount := 0
	clickToZeroCount := 0

	for _, line := range lines {
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		helpers.Check(err)

		for i := 0; i < distance; i++ {
			if direction == 'R' {
				num++
				if num > 99 {
					num = 0
				}
				if num == 0 {
					clickToZeroCount++
				}
			} else {
				num--
				if num < 0 {
					num = 99
				}
				if num == 0 {
					clickToZeroCount++
				}
			}
		}
		if num == 0 {
			pointAtZeroCount++
		}
	}

	fmt.Printf("Point-at-Zero Password: %d\n", pointAtZeroCount)
	fmt.Printf("Click-to-Zero Password: %d\n", clickToZeroCount)
}
