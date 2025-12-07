package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/input"
)

type Beam struct {
	Count int
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	beams := make(map[int]Beam, 0)
	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == 'S' {
			beams[i] = Beam{Count: 1}
			break
		}
	}

	splits := 0
	for _, line := range lines[1:] {
		for i, beam := range beams {
			if line[i] == '^' {
				splits++

				if _, exists := beams[i-1]; exists {
					beams[i-1] = Beam{Count: beams[i-1].Count + beam.Count}
				} else {
					beams[i-1] = Beam{Count: beam.Count}
				}

				if _, exists := beams[i+1]; exists {
					beams[i+1] = Beam{Count: beams[i+1].Count + beam.Count}
				} else {
					beams[i+1] = Beam{Count: beam.Count}
				}

				delete(beams, i)
			}
		}
	}

	fmt.Printf("Number of splits: %d\n", splits)

	timelines := 0
	for _, beam := range beams {
		timelines += beam.Count
	}

	fmt.Printf("Number of timelines: %d\n", timelines)
}
