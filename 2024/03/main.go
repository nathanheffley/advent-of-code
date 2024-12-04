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
	line := strings.Join(lines, "")

	mulRegex := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don\'t\(\))`)
	matches := mulRegex.FindAll([]byte(line), -1)

	total := 0
	totalWithDos := 0
	enabled := true
	for _, bytes := range matches {
		instruction := string(bytes)
		fmt.Println(instruction)
		if instruction == "do()" {
			enabled = true
			continue
		}
		if instruction == "don't()" {
			enabled = false
			continue
		}
		numbersConcatString := instruction[4 : len(instruction)-1]
		numbersStrings := strings.Split(numbersConcatString, ",")
		firstNum, err := strconv.Atoi(numbersStrings[0])
		helpers.Check(err)
		secondNum, err := strconv.Atoi(numbersStrings[1])
		helpers.Check(err)
		result := firstNum * secondNum
		total += result
		if enabled {
			totalWithDos += result
		}
	}

	fmt.Println("Mul sum:", total)
	fmt.Println("Mul sum (with dos):", totalWithDos)
}
