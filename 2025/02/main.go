package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	line := input.ReadInputFileToLines("input.txt")[0]

	ranges := strings.Split(line, ",")

	twiceTotal := 0
	anyTotal := 0

	for _, r := range ranges {
		// Split the range into its start and end values
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}

		start, err := strconv.Atoi(parts[0])
		helpers.Check(err)
		end, err := strconv.Atoi(parts[1])
		helpers.Check(err)

		for num := start; num <= end; num++ {
			str := strconv.Itoa(num)
			mid := len(str) / 2
			firstHalf := str[0:mid]
			lastHalf := str[mid:]
			if firstHalf == lastHalf {
				twiceTotal += num
			}

			for i := 1; i <= mid; i++ {
				if len(str)%i != 0 {
					continue
				}

				pattern := str[0:i]
				matches := true
				for j := 0; j < len(str); j += i {
					if str[j:j+i] != pattern {
						matches = false
						break
					}
				}
				if matches {
					anyTotal += num
					break
				}
			}
		}
	}

	fmt.Printf("Twice Total: %d\n", twiceTotal)
	fmt.Printf("Any Total: %d\n", anyTotal)
}
