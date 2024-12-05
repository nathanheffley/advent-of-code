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

	rules := make(map[string][]string)

	updates := make([][]string, 0)

	findingRules := true
	for _, line := range lines {
		if line == "" {
			findingRules = false
			continue
		}

		if findingRules {
			parts := strings.Split(line, "|")

			if rules[parts[0]] == nil {
				rules[parts[0]] = []string{}
			}
			rules[parts[0]] = append(rules[parts[0]], parts[1])
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
				if checkRuleOrder(rules, key, valBefore) {
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
			return checkRuleOrder(rules, update[i], update[j])
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
