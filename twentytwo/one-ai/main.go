package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Elf struct {
	foods []int
}

func main() {
	// Read the input
	elves := readInput("input.txt")

	// Find the Elf with the most calories
	_, calories := findElfWithMostCalories(elves)

	// Print the total number of calories carried by the Elf with the most calories
	fmt.Println(calories)
}

func readInput(inputFile string) []Elf {
	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the input file line by line
	scanner := bufio.NewScanner(file)
	elves := []Elf{}
	var currentElf Elf
	for scanner.Scan() {
		line := scanner.Text()

		// If the line is empty, we have reached the end of the current Elf's inventory
		if line == "" {
			elves = append(elves, currentElf)
			currentElf = Elf{} // Reset the currentElf variable
			continue
		}

		// Parse the calories from the line and append it to the Elf's inventory
		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		currentElf.foods = append(currentElf.foods, calories)
	}

	// Return the slice of Elves
	return elves
}

func findElfWithMostCalories(elves []Elf) (Elf, int) {
	// Keep track of the Elf with the most calories and the total calories they are carrying
	var mostCaloricElf Elf
	mostCaloricElfCalories := 0

	// Loop through each Elf
	for _, elf := range elves {
		// Calculate the total number of calories for this Elf
		totalCalories := 0
		for _, food := range elf.foods {
			totalCalories += food
		}

		// If this Elf has more calories than the current most caloric Elf, update the most caloric Elf
		if totalCalories > mostCaloricElfCalories {
			mostCaloricElf = elf
			mostCaloricElfCalories = totalCalories
		}
	}

	// Return the Elf with the most calories and the total number of calories they are carrying
	return mostCaloricElf, mostCaloricElfCalories
}
