package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/input"
)

type pos struct {
	x, y int
}

type obj struct {
	pos         pos
	wall, crate bool
	double      bool
}

func (o obj) canMove(direction pos, warehouse map[pos]obj) bool {
	if o.wall {
		return false
	}

	if o.crate {
		return warehouse[pos{o.pos.x + direction.x, o.pos.y + direction.y}].canMove(direction, warehouse)
	}

	return true
}

func (o obj) move(direction pos, warehouse map[pos]obj) {
	if o.crate {
		warehouse[o.pos] = obj{
			pos:    o.pos,
			wall:   false,
			crate:  false,
			double: o.double,
		}
		newPos := pos{o.pos.x + direction.x, o.pos.y + direction.y}
		warehouse[newPos].move(direction, warehouse)
		warehouse[newPos] = obj{
			pos:    newPos,
			wall:   false,
			crate:  true,
			double: o.double,
		}
	}
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	width := len(lines[0])
	var height int

	var bot pos
	warehouse := make(map[pos]obj)
	instructions := make([]rune, 0)
	parsingWarehouse := true
	for y, line := range lines {
		if line == "" {
			height = y
			parsingWarehouse = false
			continue
		}
		for x, s := range line {
			if parsingWarehouse {
				warehouse[pos{x, y}] = obj{
					pos:    pos{x, y},
					wall:   s == '#',
					crate:  s == 'O',
					double: true,
				}
				if s == '@' {
					bot = pos{x, y}
				}
			} else {
				instructions = append(instructions, s)
			}
		}
	}

	// printMap(width, height, bot, warehouse)
	for _, instruction := range instructions {
		// var input string
		// fmt.Scanln(&input)
		// time.Sleep(1 * time.Second)

		// fmt.Println(string(instruction))
		switch instruction {
		case '^':
			bot = moveDir(pos{0, -1}, bot, warehouse)
		case '>':
			bot = moveDir(pos{1, 0}, bot, warehouse)
		case 'v':
			bot = moveDir(pos{0, 1}, bot, warehouse)
		case '<':
			bot = moveDir(pos{-1, 0}, bot, warehouse)
		}
		// printMap(width, height, bot, warehouse)
	}
	// printMap(width, height, bot, warehouse)

	total := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if warehouse[pos{x, y}].crate {
				total += (y * 100) + x
			}
		}
	}
	fmt.Printf("%d\n1517819 is correct\n", total)
}

func moveDir(direction pos, bot pos, warehouse map[pos]obj) pos {
	botOffset := pos{bot.x + direction.x, bot.y + direction.y}
	canMove := warehouse[botOffset].canMove(direction, warehouse)
	if !canMove {
		return bot
	}
	warehouse[botOffset].move(direction, warehouse)
	return pos{bot.x + direction.x, bot.y + direction.y}
}

func printMap(width int, height int, bot pos, warehouse map[pos]obj) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := pos{x, y}
			if p == bot {
				fmt.Print("@")
			} else if warehouse[p].wall {
				fmt.Print("#")
			} else if warehouse[p].crate {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
