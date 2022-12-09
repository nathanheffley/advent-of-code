package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	treeRowsData := input.ReadInputFileToLines("input.txt")

	treeRows := make([][]int, len(treeRowsData))
	for i := range treeRows {
		treeRows[i] = make([]int, len(treeRowsData[0]))
	}

	treeCols := make([][]int, len(treeRowsData[0]))
	for i := range treeCols {
		treeCols[i] = make([]int, len(treeRowsData))
	}

	for r, treeRow := range treeRowsData {
		for c, treeStr := range treeRow {
			tree, _ := strconv.Atoi(string(treeStr))
			treeRows[r][c] = tree
			treeCols[c][r] = tree
		}
	}

	coords := make(map[string]bool)

	for y := 1; y < len(treeRows)-1; y++ {
		treeRow := treeRows[y]

		leftMax := treeRow[0]
		for x := 1; x < len(treeRow)-1; x++ {
			tree := treeRow[x]
			if tree > leftMax {
				leftMax = tree
				coords[fmt.Sprintf("%d,%d", x, y)] = true
			}
		}

		rightMax := treeRow[len(treeRow)-1]
		for x := len(treeRow) - 2; x > 0; x-- {
			tree := treeRow[x]
			if tree > rightMax {
				rightMax = tree
				coords[fmt.Sprintf("%d,%d", x, y)] = true
			}
		}
	}

	for x := 1; x < len(treeCols)-1; x++ {
		treeCol := treeCols[x]

		topMax := treeCol[0]
		for y := 1; y < len(treeCol)-1; y++ {
			tree := treeCol[y]
			if tree > topMax {
				topMax = tree
				coords[fmt.Sprintf("%d,%d", x, y)] = true
			}
		}

		bottomMax := treeCol[len(treeCol)-1]
		for y := len(treeCol) - 2; y > 0; y-- {
			tree := treeCol[y]
			if tree > bottomMax {
				bottomMax = tree
				coords[fmt.Sprintf("%d,%d", x, y)] = true
			}
		}
	}

	outerVisible := len(treeRows)*2 + len(treeCols)*2 - 4
	innerVisible := len(coords)

	fmt.Printf("Outer Trees Visible: %d\n", outerVisible)
	fmt.Printf("Inside Trees Visible: %d\n", innerVisible)
	fmt.Printf("Total Trees Visible: %d\n\n", outerVisible+innerVisible)

	maxTreesVisible := 0
	for key := range coords {
		x, _ := strconv.Atoi(strings.Split(key, ",")[0])
		y, _ := strconv.Atoi(strings.Split(key, ",")[1])

		treeRow := treeRows[y]
		treeCol := treeCols[x]

		targetTreeHeight := treeRow[x]
		treesVisible := 1

		// Check Left
		for xx := x - 1; xx >= 0; xx-- {
			if xx == 0 {
				treesVisible *= x
				break
			}

			if treeRow[xx] >= targetTreeHeight {
				treesVisible *= x - xx
				break
			}
		}

		// Check Right
		for xx := x + 1; xx < len(treeRow); xx++ {
			if xx == len(treeRow)-1 {
				treesVisible *= len(treeRow) - 1 - x
				break
			}

			if treeRow[xx] >= targetTreeHeight {
				treesVisible *= xx - x
				break
			}
		}

		// Check Up
		for yy := y - 1; yy >= 0; yy-- {
			if yy == 0 {
				treesVisible *= y
				break
			}

			if treeCol[yy] >= targetTreeHeight {
				treesVisible *= y - yy
				break
			}
		}

		// Check Down
		for yy := y + 1; yy < len(treeCol); yy++ {
			if yy == len(treeCol)-1 {
				treesVisible *= len(treeCol) - 1 - y
				break
			}

			if treeCol[yy] >= targetTreeHeight {
				treesVisible *= yy - y
				break
			}
		}

		if treesVisible > maxTreesVisible {
			maxTreesVisible = treesVisible
		}
	}

	fmt.Printf("Max Trees Visible: %d\n", maxTreesVisible)
}
