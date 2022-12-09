package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	data := input.ReadInputFileToLines("input.txt")

	var cities []string
	connections := make(map[string]map[string]int)

	for _, route := range data {
		words := strings.Split(route, " ")
		cityA := words[0]
		cityB := words[2]

		distance, err := strconv.Atoi(words[4])
		helpers.Check(err)

		if _, ok := connections[cityA]; !ok {
			cities = append(cities, cityA)
			connections[cityA] = make(map[string]int)
		}
		connections[cityA][cityB] = distance

		if _, ok := connections[cityB]; !ok {
			cities = append(cities, cityB)
			connections[cityB] = make(map[string]int)
		}
		connections[cityB][cityA] = distance
	}

	routes := permutations(cities)

	shortestRoute := 9999999
	shortestRouteName := ""
	longestRoute := 0
	longestRouteName := ""
	for _, route := range routes {
		name := strings.Join(route, " -> ")
		total := 0
		for i := 0; i < len(route)-1; i++ {
			total += connections[route[i]][route[i+1]]
		}
		output := fmt.Sprintf("%s = %d\n", name, total)

		if total < shortestRoute {
			shortestRoute = total
			shortestRouteName = output
		}
		if total > longestRoute {
			longestRoute = total
			longestRouteName = output
		}
	}
	fmt.Println("Shortest route:")
	fmt.Println(shortestRouteName)
	fmt.Println("Longest route:")
	fmt.Print(longestRouteName)
}

func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
