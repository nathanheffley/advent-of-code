package main

import (
	"fmt"
	"strconv"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	commands := input.ReadInputFileToLines("input.txt")

	cycle := 1
	x := 1

	cycles := make(map[int]int)

	for _, command := range commands {
		if command == "noop" {
			cycles[cycle] = x
			cycle++
		}

		if command[0:4] == "addx" {
			cycles[cycle] = x
			cycles[cycle+1] = x
			cycle += 2
			num, _ := strconv.Atoi(command[5:])
			x += num
		}
	}

	partOne := 0
	for _, c := range []int{20, 60, 100, 140, 180, 220} {
		partOne += cycles[c] * c
	}
	fmt.Println(partOne)

	for i := 0; i < len(cycles); i++ {
		if i%40 == 0 {
			fmt.Println()
		}

		ii := i
		for {
			if ii >= 40 {
				ii = ii - 40
			}

			if ii < 40 {
				break
			}
		}

		if ii == cycles[i+1] || ii == cycles[i+1]-1 || ii == cycles[i+1]+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}
