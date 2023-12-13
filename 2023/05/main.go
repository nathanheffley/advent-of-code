package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	lines := input.ReadInputFileToLines("input.txt")

	currentParseLine := 3

	seedToSoilMapping, currentParseLine := parseMapping(lines, currentParseLine)
	soilToFertilizerMapping, currentParseLine := parseMapping(lines, currentParseLine)
	fertilizerToWaterMapping, currentParseLine := parseMapping(lines, currentParseLine)
	waterToLightMapping, currentParseLine := parseMapping(lines, currentParseLine)
	lightToTemperatureMapping, currentParseLine := parseMapping(lines, currentParseLine)
	temperatureToHumidityMapping, currentParseLine := parseMapping(lines, currentParseLine)
	humidityToLocationMapping, _ := parseMapping(lines, currentParseLine)

	seedString := strings.Split(lines[0], ": ")[1]
	seeds := helpers.StringToNumSlice(seedString, " ")

	part1MinSeed := math.MaxInt64
	for _, seed := range seeds {
		seed = moveSeed(seed, seedToSoilMapping)
		seed = moveSeed(seed, soilToFertilizerMapping)
		seed = moveSeed(seed, fertilizerToWaterMapping)
		seed = moveSeed(seed, waterToLightMapping)
		seed = moveSeed(seed, lightToTemperatureMapping)
		seed = moveSeed(seed, temperatureToHumidityMapping)
		seed = moveSeed(seed, humidityToLocationMapping)
		if seed < part1MinSeed {
			part1MinSeed = seed
		}
	}

	part2MinSeed := math.MaxInt64
	for i := 0; i < len(seeds); i += 2 {
		for j := 0; j < seeds[i+1]; j++ {
			seed := seeds[i] + j
			seed = moveSeed(seed, seedToSoilMapping)
			seed = moveSeed(seed, soilToFertilizerMapping)
			seed = moveSeed(seed, fertilizerToWaterMapping)
			seed = moveSeed(seed, waterToLightMapping)
			seed = moveSeed(seed, lightToTemperatureMapping)
			seed = moveSeed(seed, temperatureToHumidityMapping)
			seed = moveSeed(seed, humidityToLocationMapping)
			if seed < part2MinSeed {
				part2MinSeed = seed
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1MinSeed)
	fmt.Printf("Part 2: %d\n", part2MinSeed)
}

func moveSeed(seed int, mapping map[[2]int]int) int {
	for range mapping {
		for seedRange, delta := range mapping {
			if seed >= seedRange[0] && seed < seedRange[1] {
				return seed + delta
			}
		}
	}
	return seed
}

func parseMapping(lines []string, startLine int) (map[[2]int]int, int) {
	mapping := make(map[[2]int]int)
	for i, line := range lines[startLine:] {
		if line == "" {
			return mapping, startLine + i + 2
		}
		mappingNums := helpers.StringToNumSlice(line, " ")
		destinationNum := mappingNums[0]
		sourceNum := mappingNums[1]
		length := mappingNums[2]
		mapping[[2]int{
			sourceNum,
			sourceNum + length,
		}] = destinationNum - sourceNum
	}

	return make(map[[2]int]int), -1
}
