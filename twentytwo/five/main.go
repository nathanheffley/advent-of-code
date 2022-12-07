package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	stacks := make(map[int]string)
	stacks9001 := make(map[int]string)
	initializingStacks := true
	for _, line := range lines {
		if initializingStacks {
		Initialize:
			for i := 1; i < len(line); i += 4 {
				if line[i] == ' ' {
					continue
				}

				if string(line[i]) == "1" {
					initializingStacks = false
					break Initialize
				}

				stackIndex := ((i - 1) / 4) + 1

				stacks[stackIndex] = stacks[stackIndex] + string(line[i])
				stacks9001[stackIndex] = stacks[stackIndex] + string(line[i])
			}
			continue
		}

		if line == "" {
			continue
		}

		steps := strings.Split(line, " ")

		amount, _ := strconv.Atoi(steps[1])
		stackAIdx, _ := strconv.Atoi(steps[3])
		stackBIdx, _ := strconv.Atoi(steps[5])

		// CrateMover 9000 (Part 1)
		stackA := stacks[stackAIdx]
		stackB := stacks[stackBIdx]

		moved := stackA[:amount]
		stacks[stackAIdx] = stackA[amount:]
		movedRev := ""
		for _, m := range moved {
			movedRev = string(m) + movedRev
		}
		stacks[stackBIdx] = movedRev + stackB

		// CrateMover 9001 (Part 2)
		stackA = stacks9001[stackAIdx]
		stackB = stacks9001[stackBIdx]

		moved = stackA[:amount]
		stacks9001[stackAIdx] = stackA[amount:]
		stacks9001[stackBIdx] = moved + stackB
	}

	for i := 1; i <= len(stacks); i++ {
		fmt.Print(string(stacks[i][0]))
	}
	fmt.Println()

	for i := 1; i <= len(stacks9001); i++ {
		fmt.Print(string(stacks9001[i][0]))
	}
	fmt.Println()
}
