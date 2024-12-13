package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	totalTokens := 0
	prizesWon := 0
	for i := 0; i < len(lines); i += 4 {
		aButtonPresses, bButtonPresses, cannotWin := calculateLowestTokens(lines[i:i+3], 0)
		if cannotWin == nil {
			totalTokens += aButtonPresses*3 + bButtonPresses
			prizesWon++
		}
	}
	fmt.Printf("%d tokens can win %d prizes.\n", totalTokens, prizesWon)

	totalTokens = 0
	prizesWon = 0
	for i := 0; i < len(lines); i += 4 {
		aButtonPresses, bButtonPresses, cannotWin := calculateLowestTokens(lines[i:i+3], 10000000000000)
		if cannotWin == nil {
			totalTokens += aButtonPresses*3 + bButtonPresses
			prizesWon++
		}
	}
	fmt.Printf("%d tokens can win %d prizes (adjusted).\n", totalTokens, prizesWon)
}

func calculateLowestTokens(lines []string, adjustment int) (int, int, error) {
	velocityRegex := regexp.MustCompile(`X\+(\d+), Y\+(-?\d+)$`)
	aVelocityData := velocityRegex.FindSubmatch([]byte(lines[0]))
	bVelocityData := velocityRegex.FindSubmatch([]byte(lines[1]))

	aVelocityX, _ := strconv.Atoi(string(aVelocityData[1]))
	aVelocityY, _ := strconv.Atoi(string(aVelocityData[2]))

	bVelocityX, _ := strconv.Atoi(string(bVelocityData[1]))
	bVelocityY, _ := strconv.Atoi(string(bVelocityData[2]))

	prizeRegex := regexp.MustCompile(`X=(\d+), Y=(-?\d+)$`)
	prizeData := prizeRegex.FindSubmatch([]byte(lines[2]))
	prizeX, _ := strconv.Atoi(string(prizeData[1]))
	prizeY, _ := strconv.Atoi(string(prizeData[2]))
	prizeX += adjustment
	prizeY += adjustment

	aMovesValid := (prizeY*aVelocityX-prizeX*aVelocityY)%(aVelocityX*bVelocityY-aVelocityY*bVelocityX) == 0
	bMovesValid := (prizeX*bVelocityY-prizeY*bVelocityX)%(aVelocityX*bVelocityY-aVelocityY*bVelocityX) == 0
	if aMovesValid && bMovesValid {
		return (prizeX*bVelocityY - prizeY*bVelocityX) / (aVelocityX*bVelocityY - aVelocityY*bVelocityX),
			(prizeY*aVelocityX - prizeX*aVelocityY) / (aVelocityX*bVelocityY - aVelocityY*bVelocityX),
			nil
	}

	return 0, 0, fmt.Errorf("prize cannot be won")
}
