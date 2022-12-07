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
		children: map[string]*directory{},
		fileSize: 0,
	}

	directoryStack := []*directory{&rootDirectory}

	for _, outputLine := range outputLines[1:] {
		directoryStack = handleOutput(outputLine, directoryStack)
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

func handleOutput(outputLine string, directoryStack []*directory) []*directory {
	if outputLine[0] == '$' {
		return handleCommand(outputLine, directoryStack)
	}

	return handleItem(outputLine, directoryStack)
}

func handleCommand(outputLine string, directoryStack []*directory) []*directory {
	if outputLine[2:4] != "cd" {
		return directoryStack
	}

	cdArg := outputLine[5:]

	if cdArg == ".." {
		return directoryStack[1:]
	}

	subDirectory := directoryStack[0].children[cdArg]
	return append([]*directory{subDirectory}, directoryStack...)
}

func handleItem(outputLine string, directoryStack []*directory) []*directory {
	if outputLine[:3] == "dir" {
		newDirectoryName := outputLine[4:]
		directoryStack[0].children[newDirectoryName] = &directory{
			name:     newDirectoryName,
			children: make(map[string]*directory),
			fileSize: 0,
		}
		return directoryStack
	}

	fileSize, _ := strconv.Atoi(strings.Split(outputLine, " ")[0])
	for i := 0; i < len(directoryStack); i++ {
		directoryStack[i].fileSize += fileSize
	}
	return directoryStack
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
