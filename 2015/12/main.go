package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	// json := "[1,2,3]"
	jsonLines := input.ReadInputFileToLines("input.txt")
	json := jsonLines[0]

	re := regexp.MustCompile(`((-)?\d+)`)

	numberStrs := re.FindAllString(json, -1)

	numbers := make([]int, len(numberStrs))
	for i, v := range numberStrs {
		num, err := strconv.Atoi(v)
		numbers[i] = num
		helpers.Check(err)
	}

	fmt.Println(helpers.Sum(numbers))
}
