package main

import (
	"fmt"
	"strconv"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	twoTotal := 0
	twelveTotal := 0
	for _, line := range lines {
		twoTotal += findLargestNumber(line, 2)
		twelveTotal += findLargestNumber(line, 12)
	}

	fmt.Printf("Two Battery Joltage: %d\n", twoTotal)
	fmt.Printf("Twelve Battery Joltage: %d\n", twelveTotal)
}

func findLargestNumber(line string, length int) int {
	index := 0
	numberString := ""
	for len(numberString) < length {
		endingIndex := len(line) - length + len(numberString) + 1
		largestNumber, largestIndex := firstLargestNumber(line[index:endingIndex])
		numberString += largestNumber
		index += largestIndex + 1
	}
	number, err := strconv.Atoi(numberString)
	helpers.Check(err)
	return number
}

func firstLargestNumber(line string) (string, int) {
	index := 0
	largestNumber := 0
	for i, char := range line {
		num, err := strconv.Atoi(string(char))
		helpers.Check(err)
		if num > largestNumber {
			index = i
			largestNumber = num
		}
	}
	return strconv.Itoa(largestNumber), index
}
