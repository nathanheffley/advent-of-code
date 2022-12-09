package main

import (
	"fmt"
	"math"
)

func main() {
	num := 368078

	sqrtNum := math.Sqrt(float64(num))

	cornerNum := int(math.Ceil(sqrtNum))
	if cornerNum%2 == 0 {
		cornerNum += 1
	}

	edgeNum := cornerNum - 1

	manhattanNum := (cornerNum * cornerNum) - num

	for i := 0; manhattanNum > edgeNum; i++ {
		manhattanNum = manhattanNum - 2
		fmt.Printf("manhattanNum: %d\n", manhattanNum)
	}

	fmt.Printf("manhattanNum: %d\n", manhattanNum)

	calculateLargeNum()
}

func calculateLargeNum() {
	var nums []int

	nums = append(nums, 1, 1, 2, 4, 5, 10, 11, 23, 25, 26, 54, 57, 59, 122, 133, 142, 147, 304, 330, 351, 362, 747, 806, 880, 931, 957, 1968)

	// innerRange := 3
	centerIndex := 10
	for i := len(nums); i < 900; i++ {
		// if innerRange
		newNum := nums[i-1] + nums[centerIndex] + nums[centerIndex-1] + nums[centerIndex+1]
		nums = append(nums, newNum)
		centerIndex++
		i = 901
	}

	fmt.Println(nums)
}
