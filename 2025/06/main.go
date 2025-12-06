package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

type Problem struct {
	HorizontalNumbers []int
	VerticalNumbers   []int
	Operator          string
	Width             int
}

func main() {
	lines := input.ReadInputFileToLines("input.test.txt")

	numberLinesData := lines[:len(lines)-1]
	operatorLineData := lines[len(lines)-1]

	problems := make([]Problem, 0)

	operatorStrings := strings.Split(operatorLineData, " ")
	spaces := 0
	for i := 0; i < len(operatorStrings); i++ {
		operator := operatorStrings[i]
		for i++; i < len(operatorStrings); i++ {
			if operatorStrings[i] == "" {
				spaces++
			} else {
				i--
				break
			}
		}
		problems = append(problems, Problem{
			Operator: operator,
			Width:    spaces + 1,
		})
		spaces = 0
	}

	offset := 0
	for i := 0; i < len(problems); i++ {
		numbers := make([]string, 0)
		for j := 0; j < len(numberLinesData); j++ {
			numbers = append(numbers, numberLinesData[j][offset:offset+problems[i].Width])
		}

		// Horizontal
		for j := 0; j < len(numberLinesData); j++ {
			num, err := strconv.Atoi(strings.Trim(numbers[j], " "))
			helpers.Check(err)
			problems[i].HorizontalNumbers = append(problems[i].HorizontalNumbers, num)
		}

		// Vertical
		for j := 0; j < problems[i].Width; j++ {
			numStr := ""
			for k := 0; k < len(numberLinesData); k++ {
				numStr += string(numbers[k][j])
			}
			num, err := strconv.Atoi(strings.Trim(numStr, " "))
			helpers.Check(err)
			problems[i].VerticalNumbers = append(problems[i].VerticalNumbers, num)
		}

		offset += problems[i].Width + 1
	}

	horizontalTotal := 0
	verticalTotal := 0
	for _, problem := range problems {
		var horizontalResult int
		var verticalResult int
		switch problem.Operator {
		case "+":
			horizontalResult = helpers.SumSlice(problem.HorizontalNumbers)
			verticalResult = helpers.SumSlice(problem.VerticalNumbers)
		case "*":
			horizontalResult = helpers.MultSlice(problem.HorizontalNumbers)
			verticalResult = helpers.MultSlice(problem.VerticalNumbers)
		default:
			fmt.Println("Unknown operator:", problem.Operator)
		}
		horizontalTotal += horizontalResult
		verticalTotal += verticalResult
	}

	fmt.Printf("Horizontal Total: %d\n", horizontalTotal)
	fmt.Printf("Vertical Total: %d\n", verticalTotal)
}
