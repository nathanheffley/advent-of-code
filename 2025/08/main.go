package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

type Coords struct {
	X int
	Y int
	Z int
}

func (c Coords) DistanceTo(other Coords) int {
	sums := math.Pow(float64(c.X-other.X), 2) + math.Pow(float64(c.Y-other.Y), 2) + math.Pow(float64(c.Z-other.Z), 2)
	return int(math.Sqrt(sums))
}

type JunctionBox struct {
	Id        int
	Coords    Coords
	CircuitId int
}

type Circuit struct {
	Id    int
	Boxes []JunctionBox
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	circuits := make(map[int]Circuit, len(lines))
	for i, line := range lines {
		coordData := strings.Split(line, ",")
		var x, y, z int
		for j, coord := range coordData {
			c, err := strconv.Atoi(coord)
			helpers.Check(err)
			switch j {
			case 0:
				x = c
			case 1:
				y = c
			case 2:
				z = c
			}
		}
		coords := Coords{X: x, Y: y, Z: z}
		circuits[i] = Circuit{
			Id: i,
			Boxes: []JunctionBox{
				{
					Id:        i,
					Coords:    coords,
					CircuitId: i,
				},
			},
		}
	}

	pairedBoxes := make(map[int]map[int]bool, 0)

	for i := 0; i < 1000; i++ {
		boxes := make([]JunctionBox, 0)
		for _, circuit := range circuits {
			boxes = append(boxes, circuit.Boxes...)
		}
		minDistance := math.MaxInt
		var boxA, boxB JunctionBox
		for j := 0; j < len(boxes); j++ {
			for k := j + 1; k < len(boxes); k++ {
				if boxes[j].Id == boxes[k].Id {
					continue
				}
				if boxes[j].Id < boxes[k].Id {
					if pairedBoxes[boxes[j].Id][boxes[k].Id] {
						continue
					}
				} else {
					if pairedBoxes[boxes[k].Id][boxes[j].Id] {
						continue
					}
				}
				distance := boxes[j].Coords.DistanceTo(boxes[k].Coords)
				if distance < minDistance {
					minDistance = distance
					if boxes[j].Id < boxes[k].Id {
						boxA = boxes[j]
						boxB = boxes[k]
					} else {
						boxA = boxes[k]
						boxB = boxes[j]
					}
				}
			}
		}
		_, exists := pairedBoxes[boxA.Id]
		if !exists {
			pairedBoxes[boxA.Id] = make(map[int]bool, 0)
		}
		pairedBoxes[boxA.Id][boxB.Id] = true

		if boxA.CircuitId == boxB.CircuitId {
			continue
		}

		circuitA := circuits[boxA.CircuitId]
		circuitB := circuits[boxB.CircuitId]

		updatedCircuitBBoxes := make([]JunctionBox, len(circuitB.Boxes))
		for j, box := range circuitB.Boxes {
			box.CircuitId = circuitA.Id
			updatedCircuitBBoxes[j] = box
		}

		circuitA.Boxes = append(circuitA.Boxes, updatedCircuitBBoxes...)
		circuits[circuitA.Id] = circuitA
		delete(circuits, circuitB.Id)
	}

	circuitSizes := make([]int, len(circuits))
	for _, circuit := range circuits {
		circuitSizes = append(circuitSizes, len(circuit.Boxes))
	}
	sort.Slice(circuitSizes, func(i, j int) bool {
		return circuitSizes[i] > circuitSizes[j]
	})
	fmt.Println("Mult of 3 largest circuits after 1k connections:", helpers.MultSlice(circuitSizes[:3]))

	for {
		boxes := make([]JunctionBox, 0)
		for _, circuit := range circuits {
			boxes = append(boxes, circuit.Boxes...)
		}
		minDistance := math.MaxInt
		var boxA, boxB JunctionBox
		for j := 0; j < len(boxes); j++ {
			for k := j + 1; k < len(boxes); k++ {
				if boxes[j].Id == boxes[k].Id {
					continue
				}
				if boxes[j].CircuitId == boxes[k].CircuitId {
					continue
				}
				if boxes[j].Id < boxes[k].Id {
					if pairedBoxes[boxes[j].Id][boxes[k].Id] {
						continue
					}
				} else {
					if pairedBoxes[boxes[k].Id][boxes[j].Id] {
						continue
					}
				}
				distance := boxes[j].Coords.DistanceTo(boxes[k].Coords)
				if distance < minDistance {
					minDistance = distance
					if boxes[j].Id < boxes[k].Id {
						boxA = boxes[j]
						boxB = boxes[k]
					} else {
						boxA = boxes[k]
						boxB = boxes[j]
					}
				}
			}
		}
		_, exists := pairedBoxes[boxA.Id]
		if !exists {
			pairedBoxes[boxA.Id] = make(map[int]bool, 0)
		}
		pairedBoxes[boxA.Id][boxB.Id] = true

		if boxA.CircuitId == boxB.CircuitId {
			continue
		}

		circuitA := circuits[boxA.CircuitId]
		circuitB := circuits[boxB.CircuitId]

		updatedCircuitBBoxes := make([]JunctionBox, len(circuitB.Boxes))
		for j, box := range circuitB.Boxes {
			box.CircuitId = circuitA.Id
			updatedCircuitBBoxes[j] = box
		}

		circuitA.Boxes = append(circuitA.Boxes, updatedCircuitBBoxes...)
		circuits[circuitA.Id] = circuitA
		delete(circuits, circuitB.Id)

		if len(circuits) <= 1 {
			answer := boxA.Coords.X * boxB.Coords.X
			fmt.Println("Final connection X mult:", answer)
			break
		}
	}
}
