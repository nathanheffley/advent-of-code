package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	keyBeforeValuesRules := make(map[string][]string)
	keyAfterValuesRules := make(map[string][]string)

	updates := make([][]string, 0)

	findingRules := true
	for _, line := range lines {
		if line == "" {
			findingRules = false
			continue
		}

		if findingRules {
			parts := strings.Split(line, "|")

			if keyBeforeValuesRules[parts[0]] == nil {
				keyBeforeValuesRules[parts[0]] = []string{}
			}
			keyBeforeValuesRules[parts[0]] = append(keyBeforeValuesRules[parts[0]], parts[1])

			if keyAfterValuesRules[parts[1]] == nil {
				keyAfterValuesRules[parts[1]] = []string{}
			}
			keyAfterValuesRules[parts[1]] = append(keyAfterValuesRules[parts[1]], parts[0])
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	validUpdates := make([][]string, 0)
	invalidUpdates := make([][]string, 0)

checkUpdate:
	for _, update := range updates {
		for i, key := range update {
			for _, valBefore := range update[:i] {
				if checkRuleOrder(keyBeforeValuesRules, key, valBefore) {
					invalidUpdates = append(invalidUpdates, update)
					continue checkUpdate
				}
			}

			for _, valAfter := range update[i+1:] {
				if checkRuleOrder(keyAfterValuesRules, key, valAfter) {
					invalidUpdates = append(invalidUpdates, update)
					continue checkUpdate
				}
			}
		}
		validUpdates = append(validUpdates, update)
	}

	validAnswer := 0
	for _, update := range validUpdates {
		validAnswer += getMiddleNumber(update)
	}

	invalidAnswer := 0
	for _, update := range invalidUpdates {
		sort.Slice(update, func(i, j int) bool {
			return checkRuleOrder(keyBeforeValuesRules, update[i], update[j])
		})
		invalidAnswer += getMiddleNumber(update)
	}

	fmt.Printf("validAnswer: %v\n", validAnswer)
	fmt.Printf("invalidAnswer: %v\n", invalidAnswer)
}

func checkRuleOrder(rules map[string][]string, key string, value string) bool {
	return rules[key] != nil && slices.Contains(rules[key], value)
}

func getMiddleNumber(update []string) int {
	num, err := strconv.Atoi(update[(len(update)-1)/2])
	helpers.Check(err)
	return num
}
