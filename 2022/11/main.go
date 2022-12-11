package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/input"
)

type monkey struct {
	items     []int
	operation func(item int) int
	test      int
	opTrue    int
	opFalse   int
	inspected int
}

func main() {
	data := input.ReadInputFileToLines("input.txt")

	monkeys20 := make(map[int]*monkey)
	monkeys10000 := make(map[int]*monkey)
	monkeyInt := 0
	for i := 0; i < len(data); i += 7 {
		itemStrs := strings.Split(data[i+1][18:], ", ")
		items := make([]int, len(itemStrs))
		for ii, item := range itemStrs {
			items[ii], _ = strconv.Atoi(item)
		}

		operationStr := data[i+2][23:]
		operationBits := strings.Split(operationStr, " ")
		operation := func(item int) int {
			operator := operationBits[0]
			secondBit := operationBits[1]
			var second int

			if secondBit == "old" {
				second = item
			} else {
				second, _ = strconv.Atoi(secondBit)
			}

			if operator == "+" {
				return item + second
			}
			if operator == "*" {
				return item * second
			}
			panic("Uh oh, unknown operator")
		}

		test, _ := strconv.Atoi(data[i+3][21:])

		opTrue, _ := strconv.Atoi(data[i+4][29:])
		opFalse, _ := strconv.Atoi(data[i+5][30:])

		inspected := 0

		monkeys20[monkeyInt] = &monkey{
			items,
			operation,
			test,
			opTrue,
			opFalse,
			inspected,
		}
		monkeys10000[monkeyInt] = &monkey{
			items,
			operation,
			test,
			opTrue,
			opFalse,
			inspected,
		}
		monkeyInt++
	}

	evaluateMonkeys(monkeys20, 20, true)
	fmt.Printf("Part 1: %d\n", findMonkeyBusiness(monkeys20))

	evaluateMonkeys(monkeys10000, 10000, false)
	fmt.Printf("Part 2: %d\n", findMonkeyBusiness(monkeys10000))
}

func evaluateMonkey(m *monkey, monkeys map[int]*monkey, divideFear bool, modFear int) {
	m.inspected += len(m.items)
	for _, item := range m.items {
		newItem := m.operation(item)

		if divideFear {
			newItem = newItem / 3
		} else {
			newItem = newItem % modFear
		}

		if newItem%m.test == 0 {
			monkeys[m.opTrue].items = append(monkeys[m.opTrue].items, newItem)
		} else {
			monkeys[m.opFalse].items = append(monkeys[m.opFalse].items, newItem)
		}
	}
	m.items = []int{}
}

func evaluateMonkeys(monkeys map[int]*monkey, rounds int, divideFear bool) {
	modFear := -1
	if !divideFear {
		modFear = 1
		for _, m := range monkeys {
			modFear *= m.test
		}
	}

	for i := 0; i < rounds; i++ {
		for im := 0; im < len(monkeys); im++ {
			evaluateMonkey(monkeys[im], monkeys, divideFear, modFear)
		}
	}
}

func findMonkeyBusiness(monkeys map[int]*monkey) int {
	highest := 0
	secondHighest := 0

	for im := 0; im < len(monkeys); im++ {
		if monkeys[im].inspected > highest {
			secondHighest = highest
			highest = monkeys[im].inspected
		} else if monkeys[im].inspected > secondHighest {
			secondHighest = monkeys[im].inspected
		}
	}

	return highest * secondHighest
}
