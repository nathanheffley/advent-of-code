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
	lines := input.ReadInputFileToLines("input.txt")

	part1Nums := make([]int, 0)
	part2Nums := make([]int, 0)
	for _, line := range lines {
		part1Regex := regexp.MustCompile("[1-9]")
		part1Indices := part1Regex.FindAllStringIndex(line, -1)
		part1FirstNumStr := condenseNumber(line[part1Indices[0][0]:part1Indices[0][1]])
		part1LastNumStr := condenseNumber(line[part1Indices[len(part1Indices)-1][0]:part1Indices[len(part1Indices)-1][1]])
		part1Num, err := strconv.Atoi(fmt.Sprintf("%s%s", part1FirstNumStr, part1LastNumStr))
		helpers.Check(err)
		part1Nums = append(part1Nums, part1Num)

		newLine := strings.ReplaceAll(line, "one", "o1e")
		newLine = strings.ReplaceAll(newLine, "two", "t2o")
		newLine = strings.ReplaceAll(newLine, "three", "t3e")
		newLine = strings.ReplaceAll(newLine, "four", "f4r")
		newLine = strings.ReplaceAll(newLine, "five", "f5e")
		newLine = strings.ReplaceAll(newLine, "six", "s6x")
		newLine = strings.ReplaceAll(newLine, "seven", "s7n")
		newLine = strings.ReplaceAll(newLine, "eight", "e8t")
		newLine = strings.ReplaceAll(newLine, "nine", "n9e")

		part2Regex := regexp.MustCompile("([1-9]|one|two|three|four|five|six|seven|eight|nine)")
		part2Indices := part2Regex.FindAllStringIndex(newLine, -1)
		part2FirstNumStr := condenseNumber(newLine[part2Indices[0][0]:part2Indices[0][1]])
		part2LastNumStr := condenseNumber(newLine[part2Indices[len(part2Indices)-1][0]:part2Indices[len(part2Indices)-1][1]])
		part2Num, err := strconv.Atoi(fmt.Sprintf("%s%s", part2FirstNumStr, part2LastNumStr))
		helpers.Check(err)
		part2Nums = append(part2Nums, part2Num)
	}

	fmt.Printf("Part 1 Sum: %d\n", helpers.Sum(part1Nums))
	fmt.Printf("Part 2 Sum: %d\n", helpers.Sum(part2Nums))
}

func condenseNumber(num string) string {
	if num == "one" {
		return "1"
	} else if num == "two" {
		return "2"
	} else if num == "three" {
		return "3"
	} else if num == "four" {
		return "4"
	} else if num == "five" {
		return "5"
	} else if num == "six" {
		return "6"
	} else if num == "seven" {
		return "7"
	} else if num == "eight" {
		return "8"
	} else if num == "nine" {
		return "9"
	}
	return num
}
