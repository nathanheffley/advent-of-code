package main

import (
	"slices"
	"testing"
)

var isvalidtests = []struct {
	line   string
	checks []int
	want   bool
}{
	{"#.#.###", []int{1, 1, 3}, true},
	{".#...#....###.", []int{1, 1, 3}, true},
	{".#.###.#.######", []int{1, 3, 1, 6}, true},
	{"####.#...#...", []int{4, 1, 1}, true},
	{"#....######..#####.", []int{1, 6, 5}, true},
	{".###.##....#", []int{3, 2, 1}, true},

	{"#.#.###", []int{1, 2, 3}, false},
	{".#...#....###.", []int{1, 1, 1, 3}, false},
	{".#.###.#.######", []int{1, 3, 6}, false},
	{"####.#...#...", []int{4}, false},
	{"#....######..#####.", []int{1, 5, 5}, false},
	{".###.##....#", []int{2, 2, 1}, false},
	{"........", []int{1}, false},
	{".###..##.#.#", []int{3, 2, 1}, false},
}

func TestIsValid(t *testing.T) {
	for _, tt := range isvalidtests {
		result := IsValid(tt.line, tt.checks)
		if result != tt.want {
			t.Fatalf(`IsValid("%s", %d) = %t, want %t`, tt.line, tt.checks, result, tt.want)
		}
	}
}

var currentgroupstests = []struct {
	line string
	want []int
}{
	{"#.#.###", []int{1, 1}},
	{".#...#....###.", []int{1, 1, 3}},
	{".#.###.#.######", []int{1, 3, 1}},
	{"####.#...#...", []int{4, 1, 1}},
	{"#....######..#####.", []int{1, 6, 5}},
	{".###.##....#", []int{3, 2}},
}

func TestCurrentGroups(t *testing.T) {
	for _, tt := range currentgroupstests {
		result := CurrentGroups(tt.line)
		if len(result) != len(tt.want) {
			t.Fatalf(`CurrentGroups("%s") = %v, want %v`, tt.line, result, tt.want)
		}
		for _, v := range result {
			if !slices.Contains(tt.want, v) {
				t.Fatalf(`CurrentGroups("%s") = %v, want %v`, tt.line, result, tt.want)
			}
		}
	}
}
