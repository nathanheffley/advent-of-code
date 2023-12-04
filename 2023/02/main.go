package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	games := input.ReadInputFileToLines("input.txt")

	possibleGameIds := make([]int, 0)
	totalPower := 0
	for _, game := range games {
		idRegex := regexp.MustCompile(`\s(\d+):`)
		idString := idRegex.FindStringSubmatch(game)[1]
		id, err := strconv.Atoi(idString)
		helpers.Check(err)

		drawRegex := regexp.MustCompile(`\s[0-9a-z,\s]+`)
		draws := drawRegex.FindAllString(game, -1)
		maxR, maxG, maxB := 0, 0, 0
		for _, draw := range draws[1:] {
			for _, cubeCounts := range strings.Split(draw[1:], ", ") {
				data := strings.Split(cubeCounts, " ")
				count, err := strconv.Atoi(data[0])
				helpers.Check(err)
				if data[1] == "red" && count > maxR {
					maxR = count
				} else if data[1] == "green" && count > maxG {
					maxG = count
				} else if data[1] == "blue" && count > maxB {
					maxB = count
				}
			}
		}
		if maxR <= 12 && maxG <= 13 && maxB <= 14 {
			possibleGameIds = append(possibleGameIds, id)
		}
		power := maxR * maxG * maxB
		totalPower += power
	}

	fmt.Printf("Part 1 Total: %d\n", helpers.SumSlice(possibleGameIds))
	fmt.Printf("Part 2 Total: %d\n", totalPower)
}
