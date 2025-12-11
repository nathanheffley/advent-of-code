package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

type Machine struct {
	Schematic uint16
	Lights    uint16
	Buttons   []uint16
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	schematicRegex := regexp.MustCompile(`^\[([\.\#]+)\].+`)
	buttonRegex := regexp.MustCompile(`\(([0-9,]+)\)`)

	machines := make([]Machine, 0)
	for _, line := range lines {
		machine := Machine{
			Schematic: 0,
			Lights:    0,
			Buttons:   make([]uint16, 0),
		}

		schematicData := schematicRegex.FindStringSubmatch(line)[1]
		for s, char := range schematicData {
			if char == '#' {
				machine.Schematic |= 1 << s
			}
		}

		buttonsData := buttonRegex.FindAllStringSubmatch(line, 1000)
		for _, btnData := range buttonsData {
			btnBitsValue := 0
			btnValues := make([]int, 0)
			btnStrings := strings.Split(btnData[1], ",")
			for _, btnStr := range btnStrings {
				btnIndex, err := strconv.Atoi(btnStr)
				helpers.Check(err)

				btnBitsValue |= 1 << btnIndex
				btnValues = append(btnValues, btnIndex)
			}
			machine.Buttons = append(machine.Buttons, uint16(btnBitsValue))
		}

		machines = append(machines, machine)
	}

	pressesToSolveLights := 0
	for _, machine := range machines {
		state := machine.Lights
		steps := 0
		visitedStates := make(map[uint16]bool)
		queue := []struct {
			State uint16
			Steps int
		}{{State: state, Steps: 0}}
		solved := false

		nextSteps := func(state uint16) []uint16 {
			nextStates := make([]uint16, len(machine.Buttons))
			for i, button := range machine.Buttons {
				nextState := state ^ button
				nextStates[i] = nextState
			}
			return nextStates
		}

		for len(queue) > 0 && !solved {
			current := queue[0]
			queue = queue[1:]

			if current.State == machine.Schematic {
				steps = current.Steps
				solved = true
				break
			}

			// No need to revisit states that we have already seen,
			// they can't possibly lead to the correct answer.
			if visitedStates[current.State] {
				continue
			}
			visitedStates[current.State] = true

			for _, nextState := range nextSteps(current.State) {
				if !visitedStates[nextState] {
					queue = append(queue, struct {
						State uint16
						Steps int
					}{State: nextState, Steps: current.Steps + 1})
				}
			}
		}

		if !solved {
			panic("could not solve machine")
		}

		pressesToSolveLights += steps
	}

	fmt.Println("Presses to solve lights:", pressesToSolveLights)
}
