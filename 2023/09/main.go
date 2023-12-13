package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")
	var inputs [][]int
	for _, line := range lines {
		inputs = append(inputs, helpers.StringToNumSlice(line, " "))
	}

	part1Total := 0
	part2Total := 0
	for _, input := range inputs {
		var layers [][]int
		addLayer := func(in []int) {
			var layer []int
			for i, num := range in[:len(in)-1] {
				diff := in[i+1] - num
				layer = append(layer, diff)
			}
			layers = append(layers, layer)
		}
		addLayer(input)

		currentLayer := 0
		for notAllZeroes(layers[currentLayer]) {
			addLayer(layers[currentLayer])
			currentLayer++
		}

		addedVal := 0
		for i := len(layers) - 2; i >= 0; i-- {
			addedVal = layers[i][len(layers[i])-1] + addedVal
		}
		part1Total += input[len(input)-1] + addedVal

		preceedingVal := 0
		for i := len(layers) - 2; i >= 0; i-- {
			preceedingVal = layers[i][0] - preceedingVal
			layers[i] = append([]int{preceedingVal}, layers[i]...)
		}
		part2Total += input[0] - preceedingVal
	}

	fmt.Printf("Part 1 Total: %d\n", part1Total)
	fmt.Printf("Part 2 Total: %d\n", part2Total)
}

func notAllZeroes(input []int) bool {
	notAllZeroes := false
	for _, i := range input {
		if i != 0 {
			notAllZeroes = true
			break
		}
	}
	return notAllZeroes
}
