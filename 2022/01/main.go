package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	var elves []int
	var elf []int
	for _, line := range lines {
		if line == `` {
			elves = appendElf(elves, elf)
			elf = []int{}
			continue
		}

		num, err := strconv.Atoi(line)
		helpers.Check(err)
		elf = append(elf, num)
	}
	elves = appendElf(elves, elf)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] >= elves[j]
	})
	fmt.Printf("Max: %d\n", elves[0])
	fmt.Printf("Top Three\n")
	fmt.Printf("%d\n", elves[0])
	fmt.Printf("%d\n", elves[1])
	fmt.Printf("%d\n", elves[2])
	fmt.Printf("Backup amount: %d\n", elves[0]+elves[1]+elves[2])
}

func appendElf(elves []int, elf []int) []int {
	total := 0
	for _, v := range elf {
		total += v
	}
	elves = append(elves, total)
	return elves
}
