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

type Coord struct {
	X int
	Y int
}

func (c Coord) Equals(other Coord) bool {
	return c.X == other.X && c.Y == other.Y
}

type Rectangle struct {
	A          Coord
	B          Coord
	AreaCached int
}

func (r Rectangle) Area() int {
	width := math.Abs(float64(r.A.X-r.B.X)) + 1
	height := math.Abs(float64(r.A.Y-r.B.Y)) + 1
	return int(width * height)
}

func (r Rectangle) Crosses(edges []Edge) bool {
	rectEdges := []Edge{
		{
			A: Coord{
				X: r.A.X,
				Y: r.A.Y,
			},
			B: Coord{
				X: r.B.X,
				Y: r.A.Y,
			},
		},
		{
			A: Coord{
				X: r.B.X,
				Y: r.A.Y,
			},
			B: Coord{
				X: r.B.X,
				Y: r.B.Y,
			},
		},
		{
			A: Coord{
				X: r.B.X,
				Y: r.B.Y,
			},
			B: Coord{
				X: r.A.X,
				Y: r.B.Y,
			},
		},
		{
			A: Coord{
				X: r.A.X,
				Y: r.B.Y,
			},
			B: Coord{
				X: r.A.X,
				Y: r.A.Y,
			},
		},
	}

	for _, rectEdge := range rectEdges {
		for _, edge := range edges {
			if rectEdge.Intersects(edge) {
				return true
			}
		}
	}

	return false
}

type Edge struct {
	A Coord
	B Coord
}

func (e Edge) Intersects(o Edge) bool {
	if e.A.X == e.B.X && o.A.X == o.B.X {
		// Check to make sure they don't overlap and extend past each other
		if e.A.X != o.A.X {
			return false
		}
		if e.A.Y <= o.A.Y && e.B.Y <= o.A.Y {
			return false
		}
		if e.A.Y >= o.A.Y && e.B.Y >= o.A.Y {
			return false
		}
		return true
	}
	if e.A.Y == e.B.Y && o.A.Y == o.B.Y {
		// Check to make sure they don't overlap and extend past each other
		if e.A.Y != o.A.Y {
			return false
		}
		if e.A.X <= o.A.X && e.B.X <= o.A.X {
			return false
		}
		if e.A.X >= o.A.X && e.B.X >= o.A.X {
			return false
		}
		return true
	}

	var verticalEdge, horizontalEdge Edge
	if e.A.X == e.B.X {
		verticalEdge = e
		horizontalEdge = o
	} else {
		verticalEdge = o
		horizontalEdge = e
	}

	if horizontalEdge.A.X <= verticalEdge.A.X && horizontalEdge.B.X <= verticalEdge.B.X {
		return false
	}

	if horizontalEdge.A.X >= verticalEdge.A.X && horizontalEdge.B.X >= verticalEdge.B.X {
		return false
	}

	if verticalEdge.A.Y <= horizontalEdge.A.Y && verticalEdge.B.Y <= horizontalEdge.B.Y {
		return false
	}

	if verticalEdge.A.Y >= horizontalEdge.A.Y && verticalEdge.B.Y >= horizontalEdge.B.Y {
		return false
	}

	return true
}

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	tiles := make([]Coord, len(lines))
	edges := make([]Edge, 0)
	for i, line := range lines {
		coordData := strings.Split(line, ",")
		x, err := strconv.Atoi(coordData[0])
		helpers.Check(err)
		y, err := strconv.Atoi(coordData[1])
		helpers.Check(err)
		coords := Coord{X: x, Y: y}
		tiles[i] = coords

		if i > 0 {
			prev := tiles[i-1]
			edges = append(edges, Edge{A: prev, B: coords})
		}
	}
	edges = append(edges, Edge{A: tiles[len(tiles)-1], B: tiles[0]})

	rectangles := make([]Rectangle, 0)
	for _, tileA := range tiles {
		for _, tileB := range tiles {
			if tileA.Equals(tileB) {
				continue
			}

			rect := Rectangle{A: tileA, B: tileB}
			area := rect.Area()
			rect.AreaCached = area
			rectangles = append(rectangles, rect)
		}
	}
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].AreaCached > rectangles[j].AreaCached
	})

	fmt.Printf("Largest rectangle: %d\n", rectangles[0].AreaCached)

	largestContainedRectangle := 0
	for _, rect := range rectangles {
		if !rect.Crosses(edges) && rect.AreaCached > largestContainedRectangle {
			largestContainedRectangle = rect.AreaCached
			break
		}
	}
	fmt.Printf("Largest contained rectangle: %d\n", largestContainedRectangle)
}
