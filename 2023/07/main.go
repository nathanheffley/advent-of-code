package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	cardValues := make(map[rune]int, 13)
	cardValues['2'] = 2
	cardValues['3'] = 3
	cardValues['4'] = 4
	cardValues['5'] = 5
	cardValues['6'] = 6
	cardValues['7'] = 7
	cardValues['8'] = 8
	cardValues['9'] = 9
	cardValues['T'] = 10
	cardValues['J'] = 11
	cardValues['Q'] = 12
	cardValues['K'] = 13
	cardValues['A'] = 14

	handValues := make(map[string]int, 7)
	handValues["high card"] = 0
	handValues["one pair"] = 1
	handValues["two pair"] = 2
	handValues["three of a kind"] = 3
	handValues["full house"] = 4
	handValues["four of a kind"] = 5
	handValues["five of a kind"] = 6

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

	sort.Slice(hands, func(i, j int) bool {
		// I
		iMap := make(map[rune]int)
		for _, r := range hands[i] {
			iMap[r]++
		}

		iHandType := "high card"
		iHandHasPair := false
		iHandHasThree := false
		for _, v := range iMap {
			if v == 5 {
				iHandType = "five of a kind"
				break
			}
			if v == 4 {
				iHandType = "four of a kind"
				break
			}
			if v == 3 {
				iHandHasThree = true
				continue
			}
			if v == 2 {
				if iHandHasPair {
					iHandType = "two pair"
					iHandHasPair = false
					break
				}
				iHandHasPair = true
				continue
			}
		}
		if iHandHasThree && iHandHasPair {
			iHandType = "full house"
		} else if iHandHasThree {
			iHandType = "three of a kind"
		} else if iHandHasPair {
			iHandType = "one pair"
		}
		iHandValue := handValues[iHandType]

		// J
		jMap := make(map[rune]int)
		for _, r := range hands[j] {
			jMap[r]++
		}

		jHandType := "high card"
		jHandHasPair := false
		jHandHasThree := false
		for _, v := range jMap {
			if v == 5 {
				jHandType = "five of a kind"
				break
			}
			if v == 4 {
				jHandType = "four of a kind"
				break
			}
			if v == 3 {
				jHandHasThree = true
				continue
			}
			if v == 2 {
				if jHandHasPair {
					jHandType = "two pair"
					jHandHasPair = false
					break
				}
				jHandHasPair = true
				continue
			}
		}
		if jHandHasThree && jHandHasPair {
			jHandType = "full house"
		} else if jHandHasThree {
			jHandType = "three of a kind"
		} else if jHandHasPair {
			jHandType = "one pair"
		}
		jHandValue := handValues[jHandType]

		if iHandValue != jHandValue {
			return iHandValue < jHandValue
		}

		// Tie-breaker
		for index, ir := range hands[i] {
			jr := rune(hands[j][index])
			if ir != jr {
				return cardValues[ir] < cardValues[jr]
			}
		}

		return false
	})

	var scores []int
	for i, hand := range hands {
		scores = append(scores, handBids[hand]*(i+1))
	}

	fmt.Printf("Part 1: %d\n", helpers.SumSlice(scores))
	// fmt.Printf("Part 2: %d\n", len(part2WinningTimes))
}
