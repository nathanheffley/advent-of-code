package main

import (
	"fmt"
	"strings"

	"github.com/nathanheffley/advent-of-code/input"
)

type Object struct {
	name  string
	depth int
	moons []Object
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	com := buildObject(lines, "COM", 0)

	totalOrbits := 0
	for _, moon := range com.moons {
		totalOrbits += countOrbits(moon)
	}
	fmt.Printf("Total orbits: %d\n", totalOrbits)

	pruned := pruneObject(com, "YOU", "SAN")

	checkingMoon := pruned.moons[0]
	for len(checkingMoon.moons) < 2 {
		checkingMoon = checkingMoon.moons[0]
	}
	branchingDepth := checkingMoon.depth

	you := findObject(pruned, "YOU")
	san := findObject(pruned, "SAN")

	// Adjust depth to the plant being orbited.
	youDepth := you.depth - 1
	sanDepth := san.depth - 1

	transfers := (youDepth - branchingDepth) + (sanDepth - branchingDepth)

	fmt.Printf("Transfers needed: %d\n", transfers)
}

func buildObject(lines []string, name string, depth int) Object {
	obj := Object{
		name:  name,
		depth: depth,
	}
	needle := fmt.Sprintf("%s)", name)
	for _, line := range lines {
		if strings.HasPrefix(line, needle) {
			obj.moons = append(obj.moons, buildObject(lines, strings.Split(line, ")")[1], depth+1))
		}
	}
	return obj
}

func countOrbits(obj Object) int {
	total := obj.depth
	for _, moon := range obj.moons {
		total += countOrbits(moon)
	}
	return total
}

func pruneObject(obj Object, a string, b string) Object {
	if obj.name == a || obj.name == b {
		return obj
	}
	newMoons := []Object{}
	for _, moon := range obj.moons {
		pruned := pruneObject(moon, a, b)
		if pruned.name != "" {
			newMoons = append(newMoons, pruned)
		}
	}
	obj.moons = newMoons
	if len(obj.moons) == 0 {
		return Object{}
	}
	return obj
}

func findObject(obj Object, name string) Object {
	if obj.name == name {
		return obj
	}
	for _, moon := range obj.moons {
		found := findObject(moon, name)
		if found.name != "" {
			return found
		}
	}
	return Object{}
}
