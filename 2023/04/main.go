package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	cardsData := input.ReadInputFileToLines("input.txt")

	originalCardLookup := make(map[string]string)
	cardNumberOrder := make([]string, 0)
	allCardTotals := make(map[string]int)
	for _, card := range cardsData {
		cardSplit := strings.Split(card, ": ")
		cardName := cardSplit[0]
		numbers := cardSplit[1]
		originalCardLookup[numbers] = cardName
		cardNumberOrder = append(cardNumberOrder, cardName)
		allCardTotals[cardName] = 1
	}

	cards := cardsData

	totalScore := 0
	for index, card := range cards {
		cardScore := 0
		newCardsScore := 0
		cardSplit := strings.Split(card, ": ")
		numbersJoinedString := cardSplit[1]
		numbersStrings := strings.Split(numbersJoinedString, " | ")
		winningNumbersData := numbersStrings[0]
		for _, ourNumber := range strings.Split(numbersStrings[1], " ") {
			if ourNumber == "" {
				continue
			}
			numberRegex := regexp.MustCompile(fmt.Sprintf("\\b\\s*%s\\s*\\b", ourNumber))
			if numberRegex.MatchString(winningNumbersData) {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore *= 2
				}
				newCardsScore++
			}
		}
		totalScore += cardScore
		if newCardsScore > 0 {
			cardNumber := originalCardLookup[numbersJoinedString]
			for i := 1; i <= newCardsScore; i++ {
				allCardTotals[cardNumberOrder[index+i]] += allCardTotals[cardNumber]
			}
		}
	}

	fmt.Printf("Part 1: %d\n", totalScore)
	fmt.Printf("Part 2: %d\n", helpers.SumMap(allCardTotals))
}
