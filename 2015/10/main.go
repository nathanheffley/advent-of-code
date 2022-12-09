package main

import (
	"fmt"
	"time"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	startTime := time.Now()

	value := input.ReadInputFileToLines("input.txt")[0]
	for i := 1; i <= 40; i++ {
		value = lookAndSay(value)
	}
	fmt.Println(time.Since(startTime))
	fmt.Printf("40x Length: %d\n", len(value))

	for i := 41; i <= 50; i++ {
		value = lookAndSay(value)
		fmt.Printf("Calculated iteration #%d in %s\n", i, time.Since(startTime))
	}
	fmt.Println(time.Since(startTime))
	fmt.Printf("50x Length: %d\n", len(value))
}

func lookAndSay(value string) string {
	newValue := ""
	newChunk := ""
	for i := 0; i < len(value); i++ {
		newChunk = newChunk + string(value[i])

		if i == len(value)-1 {
			newValue = newValue + fmt.Sprint(len(newChunk)) + string(newChunk[0])
			continue
		}

		if value[i] != value[i+1] {
			newValue = newValue + fmt.Sprint(len(newChunk)) + string(newChunk[0])
			newChunk = ""
			continue
		}
	}
	return newValue
}
