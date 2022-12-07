package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	stream := input.ReadInputFileToLines("input.txt")[0]

	fmt.Println(findMarkerIndex(stream, 4))

	fmt.Println(findMarkerIndex(stream, 14))
}

func findMarkerIndex(stream string, n int) int {
	for i := n; i < len(stream); i++ {
		if countUnique(stream[i-n:i]) == n {
			return i
		}
	}

	return -1
}

func countUnique(str string) int {
	allKeys := make(map[rune]bool)
	for _, r := range str {
		if _, value := allKeys[r]; !value {
			allKeys[r] = true
		}
	}
	return len(allKeys)
}
