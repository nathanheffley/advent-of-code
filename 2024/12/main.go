package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/input"
)

type point struct {
	x, y int
}

type plot struct {
	point   point
	name    rune
	grouped bool
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	garden := make(map[point]plot)
	for y, line := range lines {
		for x, char := range line {
			garden[point{x, y}] = plot{point{x, y}, char, false}
		}
	}

	perimeterCost := 0
	sidesCost := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			currentPlot := garden[point{x, y}]
			if currentPlot.grouped {
				continue
			}

			area, border, corners := checkArea(0, 0, 0, currentPlot, garden)

			fmt.Printf("%s: %d * %d\n", string(currentPlot.name), area, corners)
			perimeterCost += area * border
			sidesCost += area * corners
		}
	}

	fmt.Printf("Perimeter cost: %d\n", perimeterCost)
	fmt.Printf("Sides cost: %d\n", sidesCost)
}

func countBorders(garden map[point]plot, p plot) int {
	borders := 0
	for _, neighbor := range []point{
		{p.point.x - 1, p.point.y},
		{p.point.x + 1, p.point.y},
		{p.point.x, p.point.y - 1},
		{p.point.x, p.point.y + 1},
	} {
		if neighborPlot, ok := garden[neighbor]; !ok || neighborPlot.name != p.name {
			borders++
		}
	}
	return borders
}

func countCorners(garden map[point]plot, p plot) int {
	corners := 0

	topNeighbor, hasTopNeighbor := garden[point{p.point.x, p.point.y - 1}]
	topRightNeighbor, hasTopRightNeighbor := garden[point{p.point.x + 1, p.point.y - 1}]
	rightNeighbor, hasRightNeighbor := garden[point{p.point.x + 1, p.point.y}]
	bottomRightNeighbor, hasBottomRightNeighbor := garden[point{p.point.x + 1, p.point.y + 1}]
	bottomNeighbor, hasBottomNeighbor := garden[point{p.point.x, p.point.y + 1}]
	bottomLeftNeighbor, hasBottomLeftNeighbor := garden[point{p.point.x - 1, p.point.y + 1}]
	leftNeighbor, hasLeftNeighbor := garden[point{p.point.x - 1, p.point.y}]
	topLeftNeighbor, hasTopLeftNeighbor := garden[point{p.point.x - 1, p.point.y - 1}]

	topEdge := !hasTopNeighbor || topNeighbor.name != p.name
	rightEdge := !hasRightNeighbor || rightNeighbor.name != p.name
	bottomEdge := !hasBottomNeighbor || bottomNeighbor.name != p.name
	leftEdge := !hasLeftNeighbor || leftNeighbor.name != p.name

	if topEdge && rightEdge {
		corners++
	}
	if rightEdge && bottomEdge {
		corners++
	}
	if bottomEdge && leftEdge {
		corners++
	}
	if leftEdge && topEdge {
		corners++
	}

	if hasTopNeighbor && topNeighbor.name == p.name {
		if hasRightNeighbor && rightNeighbor.name == p.name {
			if !hasTopRightNeighbor || topRightNeighbor.name != p.name {
				corners++
			}
		}
		if hasLeftNeighbor && leftNeighbor.name == p.name {
			if !hasTopLeftNeighbor || topLeftNeighbor.name != p.name {
				corners++
			}
		}
	}

	if hasBottomNeighbor && bottomNeighbor.name == p.name {
		if hasRightNeighbor && rightNeighbor.name == p.name {
			if !hasBottomRightNeighbor || bottomRightNeighbor.name != p.name {
				corners++
			}
		}
		if hasLeftNeighbor && leftNeighbor.name == p.name {
			if !hasBottomLeftNeighbor || bottomLeftNeighbor.name != p.name {
				corners++
			}
		}
	}

	return corners
}

func checkArea(area int, border int, corners int, currentPlot plot, garden map[point]plot) (int, int, int) {
	area += 1
	border += countBorders(garden, currentPlot)
	corners += countCorners(garden, currentPlot)
	garden[currentPlot.point] = plot{currentPlot.point, currentPlot.name, true}

	// check and mark all neighbors with the same name
	// and return the area and border count
	for _, neighbor := range []point{
		{currentPlot.point.x - 1, currentPlot.point.y},
		{currentPlot.point.x + 1, currentPlot.point.y},
		{currentPlot.point.x, currentPlot.point.y - 1},
		{currentPlot.point.x, currentPlot.point.y + 1},
	} {
		if neighborPlot, ok := garden[neighbor]; ok && !neighborPlot.grouped && neighborPlot.name == currentPlot.name {
			area, border, corners = checkArea(area, border, corners, neighborPlot, garden)
		}
	}

	return area, border, corners
}
