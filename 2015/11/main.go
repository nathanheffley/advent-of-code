package main

import (
	"fmt"
	"regexp"
)

func main() {
	password := "hxbxwxba"

	password = findNextPassword(password)
	fmt.Printf("Part 1: %s\n", password)

	password = findNextPassword(password)
	fmt.Printf("Part 2: %s\n", password)
}

func findNextPassword(password string) string {
	passes := false
	for !passes {
		password = incrementPassword(password)
		password = skipConfusingLetters(password)
		passes = hasIncreasingStraight(password) && hasTwoPairs(password)
	}
	return password
}

// If we find any confusing letters, bypass them by incrementing the first one
// and setting every letter after it to 'a'.
func skipConfusingLetters(password string) string {
	confusingLettersRegex := regexp.MustCompile(`[iol]`)
	if confusingLettersRegex.MatchString(password) {
		// Convert password to byte slice
		passwordBytes := []byte(password)

		// Find first confusing letter
		firstConfusingLetterIndex := confusingLettersRegex.FindStringIndex(password)[0]

		// Increment first confusing letter
		passwordBytes[firstConfusingLetterIndex]++

		// Set all letters after first confusing letter to 'a'
		for i := firstConfusingLetterIndex + 1; i < len(passwordBytes); i++ {
			passwordBytes[i] = 'a'
		}

		// Convert byte slice back to string
		return string(passwordBytes)
	}
	return password
}

func hasIncreasingStraight(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i]+1 == password[i+1] && password[i+1]+1 == password[i+2] {
			return true
		}
	}
	return false
}

func hasTwoPairs(password string) bool {
	pairsFound := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairsFound++
			i++
		}
	}
	return pairsFound >= 2
}

func incrementPassword(password string) string {
	// Convert password to byte slice
	passwordBytes := []byte(password)

	// Increment last byte
	passwordBytes[len(passwordBytes)-1]++

	// If last byte is 'z', increment previous byte
	for i := len(passwordBytes) - 1; i >= 0; i-- {
		if passwordBytes[i] == 'z'+1 {
			passwordBytes[i] = 'a'
			if i > 0 {
				passwordBytes[i-1]++
			}
		}
	}

	// Convert byte slice back to string
	return string(passwordBytes)
}
