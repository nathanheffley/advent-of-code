package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	rollMap := make(map[int]map[int]rune)
	for lineIndex, line := range lines {
		rollMap[lineIndex] = make(map[int]rune)
		for columnIndex, char := range line {
			rollMap[lineIndex][columnIndex] = char
		}
	}

	accessibleRolls, rollMap := removeAccessibleRolls(rollMap)
	fmt.Printf("Initial Accessible Rolls: %d\n", accessibleRolls)

	removedRolls := 0
	for accessibleRolls > 0 {
		removedRolls += accessibleRolls
		accessibleRolls, rollMap = removeAccessibleRolls(rollMap)
	}
	fmt.Printf("Total Removed Rolls: %d\n", removedRolls)
}

func removeAccessibleRolls(rollMap map[int]map[int]rune) (int, map[int]map[int]rune) {
	accessibleRolls := 0
	newRollMap := make(map[int]map[int]rune)

	for lineIndex, line := range rollMap {
		newRollMap[lineIndex] = make(map[int]rune)
		for columnIndex, char := range line {
			if char == '.' {
				newRollMap[lineIndex][columnIndex] = '.'
				continue
			}

			adjacentRolls := 0
			// check up left
			if lineIndex > 0 && columnIndex > 0 {
				if rollMap[lineIndex-1][columnIndex-1] == '@' {
					adjacentRolls++
				}
			}
			// check up
			if lineIndex > 0 {
				if rollMap[lineIndex-1][columnIndex] == '@' {
					adjacentRolls++
				}
			}
			// check up right
			if lineIndex > 0 && columnIndex < len(line)-1 {
				if rollMap[lineIndex-1][columnIndex+1] == '@' {
					adjacentRolls++
				}
			}
			// check left
			if columnIndex > 0 {
				if rollMap[lineIndex][columnIndex-1] == '@' {
					adjacentRolls++
				}
			}
			// check right
			if columnIndex < len(line)-1 {
				if rollMap[lineIndex][columnIndex+1] == '@' {
					adjacentRolls++
				}
			}
			// check down left
			if lineIndex < len(rollMap)-1 && columnIndex > 0 {
				if rollMap[lineIndex+1][columnIndex-1] == '@' {
					adjacentRolls++
				}
			}
			// check down
			if lineIndex < len(rollMap)-1 {
				if rollMap[lineIndex+1][columnIndex] == '@' {
					adjacentRolls++
				}
			}
			// check down right
			if lineIndex < len(rollMap)-1 && columnIndex < len(line)-1 {
				if rollMap[lineIndex+1][columnIndex+1] == '@' {
					adjacentRolls++
				}
			}

			if adjacentRolls < 4 {
				accessibleRolls++
				newRollMap[lineIndex][columnIndex] = '.'
			} else {
				newRollMap[lineIndex][columnIndex] = '@'
			}
		}
	}
	return accessibleRolls, newRollMap
}
