package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	people := make(map[string]map[string]int)

	for _, line := range lines {
		splt := strings.Split(line, " ")
		name := splt[0]
		nextToName := splt[len(splt)-1]
		nextToName = nextToName[:len(nextToName)-1]
		if _, ok := people[name]; !ok {
			people[name] = make(map[string]int)
		}
		num, err := strconv.Atoi(splt[3])
		helpers.Check(err)
		if splt[2] == "gain" {
			people[name][nextToName] = num
		} else {
			people[name][nextToName] = -num
		}
	}

	names := []string{}
	for name := range people {
		names = append(names, name)
	}

	seatings := helpers.Permutate[string](names)

	part1Max := 0
	for _, seating := range seatings {
		total := 0
		for i := 0; i < len(seating); i++ {
			name := seating[i]
			left := seating[(i-1+len(seating))%len(seating)]
			right := seating[(i+1)%len(seating)]
			total += people[name][left]
			total += people[name][right]
		}
		if total > part1Max {
			part1Max = total
		}
	}

	people["Me"] = make(map[string]int)
	for name := range people {
		people[name]["Me"] = 0
		people["Me"][name] = 0
	}
	names = append(names, "Me")

	seatings = helpers.Permutate[string](names)

	part2Max := 0
	for _, seating := range seatings {
		total := 0
		for i := 0; i < len(seating); i++ {
			name := seating[i]
			left := seating[(i-1+len(seating))%len(seating)]
			right := seating[(i+1)%len(seating)]
			total += people[name][left]
			total += people[name][right]
		}
		if total > part2Max {
			part2Max = total
		}
	}

	fmt.Printf("Part 1: %d\n", part1Max)
	fmt.Printf("Part 2: %d\n", part2Max)
}
