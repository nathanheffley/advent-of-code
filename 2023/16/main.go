package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/input"
)

type Beam struct {
	x, y, xVelocity, yVelocity int
}

type Space struct {
	char                                                    rune
	energized, goingNorth, goingEast, goingSouth, goingWest bool
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	grid := make([][]Space, len(lines))

	for y, line := range lines {
		grid[y] = make([]Space, len(line))
		for x, char := range line {
			grid[y][x] = Space{char: char}
		}
	}

	part1Total := getTotalEnergy(Beam{
		x:         0,
		y:         0,
		xVelocity: 1,
		yVelocity: 0,
	}, grid)
	fmt.Printf("Part 1: %d\n", part1Total)

	part2Total := 0
	for y := 0; y < len(grid); y++ {
		resetGrid(grid)
		fromLeftTotal := getTotalEnergy(Beam{
			x:         0,
			y:         y,
			xVelocity: 1,
			yVelocity: 0,
		}, grid)
		if fromLeftTotal > part2Total {
			part2Total = fromLeftTotal
		}
		resetGrid(grid)
		fromRightTotal := getTotalEnergy(Beam{
			x:         len(grid[0]) - 1,
			y:         y,
			xVelocity: -1,
			yVelocity: 0,
		}, grid)
		if fromRightTotal > part2Total {
			part2Total = fromRightTotal
		}
	}
	for x := 0; x < len(grid[0]); x++ {
		resetGrid(grid)
		fromTopTotal := getTotalEnergy(Beam{
			x:         x,
			y:         0,
			xVelocity: 0,
			yVelocity: 1,
		}, grid)
		if fromTopTotal > part2Total {
			part2Total = fromTopTotal
		}
		resetGrid(grid)
		fromBottomTotal := getTotalEnergy(Beam{
			x:         x,
			y:         len(grid) - 1,
			xVelocity: 0,
			yVelocity: -1,
		}, grid)
		if fromBottomTotal > part2Total {
			part2Total = fromBottomTotal
		}
	}
	fmt.Printf("Part 2: %d\n", part2Total)

	// for _, row := range grid {
	// 	for _, space := range row {
	// 		if space.energized {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}

func resetGrid(grid [][]Space) {
	for y, row := range grid {
		for x := range row {
			grid[y][x].energized = false
			grid[y][x].goingNorth = false
			grid[y][x].goingEast = false
			grid[y][x].goingSouth = false
			grid[y][x].goingWest = false
		}
	}
}

func getTotalEnergy(startBeam Beam, grid [][]Space) int {
	total := 0

	beams := []Beam{startBeam}

	for i := 0; i < len(beams); i++ {
		beam := beams[i]
		for beam.x >= 0 && beam.x < len(grid[0]) && beam.y >= 0 && beam.y < len(grid) {
			space := &grid[beam.y][beam.x]

			if space.energized {
				if space.goingNorth && beam.yVelocity == -1 {
					break
				}
				if space.goingSouth && beam.yVelocity == 1 {
					break
				}
				if space.goingEast && beam.xVelocity == 1 {
					break
				}
				if space.goingWest && beam.xVelocity == -1 {
					break
				}
			} else {
				total++
				space.energized = true
			}

			if beam.xVelocity == 1 {
				space.goingEast = true
			} else if beam.xVelocity == -1 {
				space.goingWest = true
			} else if beam.yVelocity == 1 {
				space.goingSouth = true
			} else if beam.yVelocity == -1 {
				space.goingNorth = true
			}

			if space.char == '/' {
				beam.xVelocity, beam.yVelocity = -beam.yVelocity, -beam.xVelocity
			} else if space.char == '\\' {
				beam.xVelocity, beam.yVelocity = beam.yVelocity, beam.xVelocity
			} else if space.char == '-' && beam.yVelocity != 0 {
				beam.yVelocity = 0
				beam.xVelocity = 1
				beams = append(beams, Beam{
					x:         beam.x,
					y:         beam.y,
					xVelocity: -1,
					yVelocity: 0,
				})
			} else if space.char == '|' && beam.xVelocity != 0 {
				beam.xVelocity = 0
				beam.yVelocity = 1
				beams = append(beams, Beam{
					x:         beam.x,
					y:         beam.y,
					xVelocity: 0,
					yVelocity: -1,
				})
			}

			beam.x += beam.xVelocity
			beam.y += beam.yVelocity
		}
	}

	return total
}
