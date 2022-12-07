package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/input"
)

type directory struct {
	name     string
	parent   *directory
	children map[string]*directory
	fileSize int
}

var total = 0

var distance = math.MaxInt
var closest int

func main() {
	outputLines := input.ReadInputFileToLines("input.txt")

	rootDirectory := directory{
		name:     "/",
		parent:   nil,
		children: map[string]*directory{},
		fileSize: 0,
	}

	currentDirectory := &rootDirectory

	for _, outputLine := range outputLines[1:] {
		currentDirectory = handleOutput(outputLine, currentDirectory)
	}

	under100k(rootDirectory)
	fmt.Printf("Under 10k sum: %d\n", total)

	totalDisk := 70000000
	targetDisk := 30000000
	diskAvailable := totalDisk - rootDirectory.fileSize
	diskNeeded := targetDisk - diskAvailable
	closestToDiskNeeded(rootDirectory, diskNeeded)
	fmt.Printf("Filesize to delete: %d\n", closest)
}

func handleOutput(outputLine string, currentDirectory *directory) *directory {
	if outputLine[0] == '$' {
		return handleCommand(outputLine, currentDirectory)
	}

	handleItem(outputLine, currentDirectory)
	return currentDirectory
}

func handleCommand(outputLine string, currentDirectory *directory) *directory {
	if outputLine[2:4] != "cd" {
		return currentDirectory
	}

	cdArg := outputLine[5:]

	if cdArg == ".." {
		return currentDirectory.parent
	}

	return currentDirectory.children[cdArg]
}

func handleItem(outputLine string, currentDirectory *directory) {
	if outputLine[:3] == "dir" {
		newDirectoryName := outputLine[4:]
		currentDirectory.children[newDirectoryName] = &directory{
			name:     newDirectoryName,
			parent:   currentDirectory,
			children: make(map[string]*directory),
			fileSize: 0,
		}
	}

	fileSize, _ := strconv.Atoi(strings.Split(outputLine, " ")[0])
	for {
		currentDirectory.fileSize += fileSize
		currentDirectory = currentDirectory.parent
		if currentDirectory == nil {
			break
		}
	}
}

func under100k(dir directory) {
	if dir.fileSize <= 100000 {
		total += dir.fileSize
	}
	for _, c := range dir.children {
		under100k(*c)
	}
}

func closestToDiskNeeded(dir directory, diskNeeded int) {
	dist := dir.fileSize - diskNeeded
	if dist > 0 && dist < distance {
		distance = dir.fileSize - diskNeeded
		closest = dir.fileSize
	}
	for _, c := range dir.children {
		closestToDiskNeeded(*c, diskNeeded)
	}
}
