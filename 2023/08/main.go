package main

import (
	"fmt"
	"strings"

	"github.com/nathanheffley/advent-of-code/input"
)

type Node struct {
	left  string
	right string
}

var nodes = make(map[string]Node)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	instructions := strings.Split(lines[0], "")

	for _, node := range lines[2:] {
		nodes[node[:3]] = Node{
			left:  node[7:10],
			right: node[12:15],
		}
	}

	part1Steps := getPathSteps("AAA", instructions, true)

	part2Nodes := make(map[string]int)
	for nodeName := range nodes {
		if nodeName[2] == 'A' {
			part2Nodes[nodeName] = -1
		}
	}

	for nodeName := range part2Nodes {
		steps := getPathSteps(nodeName, instructions, false)
		part2Nodes[nodeName] = steps
	}
	part2StepValues := make([]int, 0, len(part2Nodes))
	for _, steps := range part2Nodes {
		part2StepValues = append(part2StepValues, steps)
	}
	part2Steps := getLeastCommonMultiple(part2StepValues...)

	fmt.Printf("Part 1 Total: %d\n", part1Steps)
	fmt.Printf("Part 2 Total: %d\n", part2Steps)
}

func getPathSteps(startingNodeName string, instructions []string, fullyTerminate bool) int {
	steps := 0

	var endCheck func(string) bool
	if fullyTerminate {
		endCheck = func(nodeName string) bool {
			return nodeName == "ZZZ"
		}
	} else {
		endCheck = func(nodeName string) bool {
			return nodeName[2] == 'Z'
		}
	}

	currentNodeName := startingNodeName
	currentInstructionIndex := 0
	for !endCheck(currentNodeName) {
		steps++

		currentNode := nodes[currentNodeName]
		if instructions[currentInstructionIndex] == "L" {
			currentNodeName = currentNode.left
		} else {
			currentNodeName = currentNode.right
		}

		currentInstructionIndex++
		if currentInstructionIndex >= len(instructions) {
			currentInstructionIndex = 0
		}
	}

	return steps
}

func getGreatestCommonDenominator(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func getLeastCommonMultiple(integers ...int) int {
	result := integers[0] * integers[1] / getGreatestCommonDenominator(integers[0], integers[1])
	for i := 2; i < len(integers); i++ {
		result = getLeastCommonMultiple(result, integers[i])
	}
	return result
}
