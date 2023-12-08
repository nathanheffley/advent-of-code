package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

var handValues = map[string]int{
	"high card":       0,
	"one pair":        1,
	"two pair":        2,
	"three of a kind": 3,
	"full house":      4,
	"four of a kind":  5,
	"five of a kind":  6,
}

var cardValues = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	handBids := make(map[string]int)
	var hands []string

	for _, line := range lines {
		parts := strings.Split(line, " ")

		hand := parts[0]
		bid, err := strconv.Atoi(parts[1])
		helpers.Check(err)

		handBids[hand] = bid
		hands = append(hands, hand)
	}

	// No Jokers
	sort.Slice(hands, func(i, j int) bool {
		return handSortFunction(hands[i], hands[j], getHandValue)
	})

	var part1Scores []int
	for i, hand := range hands {
		part1Scores = append(part1Scores, handBids[hand]*(i+1))
	}

	// With Jokers
	cardValues['J'] = 0

	sort.Slice(hands, func(i, j int) bool {
		return handSortFunction(hands[i], hands[j], getJokerHandValue)
	})

	var part2Scores []int
	for i, hand := range hands {
		part2Scores = append(part2Scores, handBids[hand]*(i+1))
	}

	fmt.Printf("Part 1 Total: %d\n", helpers.SumSlice(part1Scores))
	fmt.Printf("Part 2 Total: %d\n", helpers.SumSlice(part2Scores))
}

func handSortFunction(a string, b string, handValueCalculator func(hand string) int) bool {
	iHandValue := handValueCalculator(a)
	jHandValue := handValueCalculator(b)

	if iHandValue != jHandValue {
		return iHandValue < jHandValue
	}

	// Tie-breaker
	for index, ir := range a {
		jr := rune(b[index])
		if ir != jr {
			return cardValues[ir] < cardValues[jr]
		}
	}

	return false
}

func GetHandType(hand string) string {
	handMap := make(map[rune]int)
	for _, r := range hand {
		handMap[r]++
	}

	if len(handMap) == 1 {
		return "five of a kind"
	}

	if len(handMap) == 2 {
		for _, v := range handMap {
			if v == 4 {
				return "four of a kind"
			}
			if v == 3 {
				return "full house"
			}
		}
	}

	if len(handMap) == 3 {
		for _, v := range handMap {
			if v == 3 {
				return "three of a kind"
			}
		}
		return "two pair"
	}

	if len(handMap) == 4 {
		return "one pair"
	}

	return "high card"
}

func getHandValue(hand string) int {
	return handValues[GetHandType(hand)]
}

func GetJokerHandType(hand string) string {
	handMap := make(map[rune]int)
	for _, r := range hand {
		handMap[r]++
	}

	jokerCount := 0
	for c, v := range handMap {
		if c == 'J' {
			jokerCount = v
			break
		}
	}

	if len(handMap) == 1 {
		return "five of a kind"
	}

	if len(handMap) == 2 {
		if jokerCount != 0 {
			return "five of a kind"
		}

		for c, v := range handMap {
			if c == 'J' {
				continue
			}
			if v == 4 {
				return "four of a kind"
			}
			if v == 3 {
				return "full house"
			}
		}
	}

	if len(handMap) == 3 {
		if jokerCount == 3 {
			return "four of a kind"
		}
		for c, v := range handMap {
			if c == 'J' {
				continue
			}
			if v == 3 && jokerCount == 1 {
				return "four of a kind"
			}
			if v == 3 {
				return "three of a kind"
			}
			if v == 2 && jokerCount == 2 {
				return "four of a kind"
			}
			if v == 2 && jokerCount == 1 {
				return "full house"
			}
		}
		return "two pair"
	}

	if len(handMap) == 4 {
		if jokerCount != 0 {
			return "three of a kind"
		}
		return "one pair"
	}

	if len(handMap) == 5 {
		if jokerCount != 0 {
			return "one pair"
		}
	}

	return "high card"
}

func getJokerHandValue(hand string) int {
	return handValues[GetJokerHandType(hand)]
}
