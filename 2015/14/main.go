package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")
	timeLimit := 2503

	r := regexp.MustCompile("^([A-Z][a-z]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.")

	reindeers := make(map[string]map[string]int)

	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		name := matches[1]
		kms, err := strconv.Atoi(matches[2])
		helpers.Check(err)
		time, err := strconv.Atoi(matches[3])
		helpers.Check(err)
		rest, err := strconv.Atoi(matches[4])
		helpers.Check(err)

		reindeers[name] = map[string]int{
			"kms":      kms,
			"time":     time,
			"rest":     rest,
			"distance": 0,
			"points":   0,
		}
	}

	for i := 0; i < timeLimit; i++ {
		for _, stats := range reindeers {
			if i%(stats["time"]+stats["rest"]) < stats["time"] {
				stats["distance"] += stats["kms"]
			}
		}

		maxDistance := max(reindeers, "distance")

		for _, stats := range reindeers {
			if stats["distance"] == maxDistance {
				stats["points"]++
			}
		}
	}

	maxDistance := max(reindeers, "distance")
	maxPoints := max(reindeers, "points")

	fmt.Printf("Furthest Distance: %d\n", maxDistance)
	fmt.Printf("Most Points: %d\n", maxPoints)
}

func max(m map[string]map[string]int, field string) int {
	max := 0
	for _, stats := range m {
		if stats[field] > max {
			max = stats[field]
		}
	}
	return max
}
