package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/input"
)

type directory struct {
	name     string
	children map[string]directory
	fileSize int
}

var total int

var distance int
var closest int

func main() {
	commands := input.ReadInputFileToLines("input.txt")

	total = 0

	distance = 999999999999

	rootDirectory := directory{
		name:     "/",
		children: map[string]directory{},
		fileSize: 0,
	}

	treeStack := []*directory{&rootDirectory}
	currentDirectory := treeStack[0]

	for _, command := range commands[1:] {
		if command[0] == '$' {
			if command[2:4] == "cd" {
				if command == "$ cd .." {
					treeStack = treeStack[1:]
					currentDirectory = treeStack[0]
				} else {
					subDirectory := currentDirectory.children[command[5:]]
					treeStack = append([]*directory{&subDirectory}, treeStack...)
					currentDirectory = treeStack[0]
				}
			}
		} else {
			if command[:3] == "dir" {
				newDirectoryName := command[4:]
				currentDirectory.children[newDirectoryName] = directory{
					name:     newDirectoryName,
					children: make(map[string]directory),
					fileSize: 0,
				}
			} else {
				fileSize, _ := strconv.Atoi(strings.Split(command, " ")[0])
				for i := 1; i < len(treeStack); i++ {
					mappedDirectory := treeStack[i].children[treeStack[i-1].name]
					mappedDirectory.fileSize += fileSize
					treeStack[i].children[treeStack[i-1].name] = mappedDirectory
				}
				treeStack[len(treeStack)-1].fileSize += fileSize
			}
		}
	}

	under100k(rootDirectory)
	fmt.Printf("Under 10k sum: %d\n", total)

	fmt.Println()

	totalDisk := 70000000
	targetDisk := 30000000
	diskAvailable := totalDisk - rootDirectory.fileSize
	diskNeeded := targetDisk - diskAvailable
	closestToDiskNeeded(rootDirectory, diskNeeded)
	fmt.Printf("Filesize to delete: %d\n", closest)
}

func under100k(dir directory) {
	if dir.fileSize <= 100000 {
		total += dir.fileSize
	}
	for _, c := range dir.children {
		under100k(c)
	}
}

func closestToDiskNeeded(dir directory, diskNeeded int) {
	dist := dir.fileSize - diskNeeded
	if dist > 0 && dist < distance {
		distance = dir.fileSize - diskNeeded
		closest = dir.fileSize
	}
	for _, c := range dir.children {
		closestToDiskNeeded(c, diskNeeded)
	}
}
