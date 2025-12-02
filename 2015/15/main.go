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

	r := regexp.MustCompile("^([A-Z][a-z]+): capacity ([-0-9]+), durability ([-0-9]+), flavor ([-0-9]+), texture ([-0-9]+), calories ([-0-9]+)$")

	ingredients := make(map[string]map[string]int)

	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		name := matches[1]
		capacity, err := strconv.Atoi(matches[2])
		helpers.Check(err)
		durability, err := strconv.Atoi(matches[3])
		helpers.Check(err)
		flavor, err := strconv.Atoi(matches[4])
		helpers.Check(err)
		texture, err := strconv.Atoi(matches[5])
		helpers.Check(err)
		calories, err := strconv.Atoi(matches[6])
		helpers.Check(err)

		ingredients[name] = map[string]int{
			"capacity":   capacity,
			"durability": durability,
			"flavor":     flavor,
			"texture":    texture,
			"calories":   calories,
		}
	}

	amounts := make(map[string]int)
	for ingredient := range ingredients {
		amounts[ingredient] = 0
	}

	maxScore := 0
	maxCalorieScore := 0
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			if (a + b) > 100 {
				continue
			}
			for c := 0; c <= 100; c++ {
				if (a + b + c) > 100 {
					continue
				}

				amounts["Sugar"] = a
				amounts["Sprinkles"] = b
				amounts["Candy"] = c
				amounts["Chocolate"] = 100 - (a + b + c)

				capacityScore := score(amounts, ingredients, "capacity")
				durabilityScore := score(amounts, ingredients, "durability")
				flavorScore := score(amounts, ingredients, "flavor")
				textureScore := score(amounts, ingredients, "texture")
				calorieScore := score(amounts, ingredients, "calories")

				totalScore := capacityScore * durabilityScore * flavorScore * textureScore
				if totalScore > maxScore {
					maxScore = totalScore
				}

				if calorieScore == 500 && totalScore > maxCalorieScore {
					maxCalorieScore = totalScore
				}
			}
		}
	}

	fmt.Printf("%d\n", maxScore)
	fmt.Printf("%d\n", maxCalorieScore)
}

func score(amounts map[string]int, ingredients map[string]map[string]int, field string) int {
	total := 0
	for ingredient, amount := range amounts {
		total += amount * ingredients[ingredient][field]
	}

	if total < 0 {
		return 0
	}

	return total
}
