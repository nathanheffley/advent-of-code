package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	beams := make(map[int]int, 0)
	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == 'S' {
			beams[i] = 1
			break
		}
	}

	splits := 0
	for _, line := range lines[1:] {
		for i, beam := range beams {
			if line[i] == '^' {
				splits++

				if _, exists := beams[i-1]; exists {
					beams[i-1] = beams[i-1] + beam
				} else {
					beams[i-1] = beam
				}

				if _, exists := beams[i+1]; exists {
					beams[i+1] = beams[i+1] + beam
				} else {
					beams[i+1] = beam
				}

				delete(beams, i)
			}
		}
	}

	fmt.Printf("Number of splits: %d\n", splits)
	fmt.Printf("Number of timelines: %d\n", helpers.SumMap(beams))
}
