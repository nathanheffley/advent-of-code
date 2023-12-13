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

	timeLine := lines[0][5:]
	distanceLine := lines[1][9:]

	times := helpers.StringToNumSlice(timeLine, " ")
	distances := helpers.StringToNumSlice(distanceLine, " ")

	part2TimeLine := strings.ReplaceAll(timeLine, " ", "")
	part2DistanceLine := strings.ReplaceAll(distanceLine, " ", "")

	part2Time, err := strconv.Atoi(part2TimeLine)
	helpers.Check(err)
	part2Distance, err := strconv.Atoi(part2DistanceLine)
	helpers.Check(err)

	part1Total := 1
	for r := 0; r < len(times); r++ {
		time := times[r]
		distance := distances[r]
		speed := 0
		winningTimes := []int{}
		for t := 0; t < time; t++ {
			currentDistance := speed * (time - t)
			if currentDistance > distance {
				winningTimes = append(winningTimes, t)
			}
			speed++
		}
		part1Total = part1Total * len(winningTimes)
	}

	part2Speed := 0
	part2WinningTimes := []int{}
	for t := 0; t < part2Time; t++ {
		currentDistance := part2Speed * (part2Time - t)
		if currentDistance > part2Distance {
			part2WinningTimes = append(part2WinningTimes, t)
		}
		part2Speed++
	}

	fmt.Printf("Part 1: %d\n", part1Total)
	fmt.Printf("Part 2: %d\n", len(part2WinningTimes))
}
