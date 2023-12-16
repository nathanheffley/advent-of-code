package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

type LensItem struct {
	label string
	focal int
}

func main() {
	line := input.ReadInputFileToLines("input.txt")[0]

	steps := strings.Split(line, ",")

	part1Total := 0
	for _, step := range steps {
		hashValue := 0
		for _, char := range step {
			hashValue += int(char)
			hashValue *= 17
			hashValue %= 256
		}
		part1Total += hashValue
	}
	fmt.Printf("Part 1: %d\n", part1Total)

	boxes := make([][]LensItem, 256)
	for _, step := range steps {
		var label string
		var number int
		var err error
		remove := false
		if strings.Contains(step, "=") {
			label = strings.Split(step, "=")[0]
			number, err = strconv.Atoi(strings.Split(step, "=")[1])
			helpers.Check(err)
		} else {
			label = strings.Split(step, "-")[0]
			remove = true
		}
		boxNum := 0
		for _, char := range label {
			boxNum += int(char)
			boxNum *= 17
			boxNum %= 256
		}

		if remove {
			for i, lens := range boxes[boxNum] {
				if lens.label == label {
					boxes[boxNum] = append(boxes[boxNum][:i], boxes[boxNum][i+1:]...)
					break
				}
			}
		} else {
			needsLens := true
			for i, lens := range boxes[boxNum] {
				if lens.label == label {
					boxes[boxNum][i] = LensItem{
						label: label,
						focal: number,
					}
					needsLens = false
					break
				}
			}
			if needsLens {
				boxes[boxNum] = append(boxes[boxNum], LensItem{
					label: label,
					focal: number,
				})
			}
		}
	}
	part2Total := 0
	for boxNum, box := range boxes {
		if len(box) > 0 {
			for i, lens := range box {
				total := boxNum + 1
				total *= i + 1
				total *= lens.focal
				part2Total += total
			}
		}
	}
	fmt.Printf("Part 2: %d\n", part2Total)
}
