package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/helpers"
	"github.com/nathanheffley/advent-of-code/intcode"
)

func main() {
	program := "3,8,1001,8,10,8,105,1,0,0,21,38,59,84,93,110,191,272,353,434,99999,3,9,101,5,9,9,1002,9,5,9,101,5,9,9,4,9,99,3,9,1001,9,3,9,1002,9,2,9,101,4,9,9,1002,9,4,9,4,9,99,3,9,102,5,9,9,1001,9,4,9,1002,9,2,9,1001,9,5,9,102,4,9,9,4,9,99,3,9,1002,9,2,9,4,9,99,3,9,1002,9,5,9,101,4,9,9,102,2,9,9,4,9,99,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,99"

	maxSignal := 0

	for _, sequence := range helpers.Permutate([]int{0, 1, 2, 3, 4}) {
		// A
		input := make(chan int)
		output := make(chan int)
		go intcode.Execute(program, input, output)
		input <- sequence[0]
		input <- 0
		result := <-output

		// B
		input = make(chan int)
		output = make(chan int)
		go intcode.Execute(program, input, output)
		input <- sequence[1]
		input <- result
		result = <-output

		// C
		input = make(chan int)
		output = make(chan int)
		go intcode.Execute(program, input, output)
		input <- sequence[2]
		input <- result
		result = <-output

		// D
		input = make(chan int)
		output = make(chan int)
		go intcode.Execute(program, input, output)
		input <- sequence[3]
		input <- result
		result = <-output

		// E
		input = make(chan int)
		output = make(chan int)
		go intcode.Execute(program, input, output)
		input <- sequence[4]
		input <- result
		result = <-output

		if result > maxSignal {
			maxSignal = result
		}
	}

	fmt.Println(maxSignal)
	fmt.Println("(official answer for part 1: 225056)")

	maxSignal = 0
	for _, sequence := range helpers.Permutate([]int{5, 6, 7, 8, 9}) {
		programResult := 0
		panicAnswer(program, sequence, &programResult)
		if programResult > maxSignal {
			maxSignal = programResult
		}
	}
	fmt.Println(maxSignal)
	fmt.Println("(official answer for part 2: 14260332)")
}

func panicAnswer(program string, sequence []int, final *int) {
	maxSignal := 0

	// Eventually, the loops below will panic when you try to send the result
	// to A but you've reached the end of the recursion. Channels are hard,
	// so it's a lot easier to just catch the inevitable panic and write out
	// the answer like this. YOLO.
	defer func() {
		if err := recover(); err != nil {
			*final = maxSignal
		}
	}()

	// A
	inputA := make(chan int)
	outputA := make(chan int)
	go intcode.Execute(program, inputA, outputA)
	inputA <- sequence[0]
	inputA <- 0
	result := <-outputA

	// B
	inputB := make(chan int)
	outputB := make(chan int)
	go intcode.Execute(program, inputB, outputB)
	inputB <- sequence[1]
	inputB <- result
	result = <-outputB

	// C
	inputC := make(chan int)
	outputC := make(chan int)
	go intcode.Execute(program, inputC, outputC)
	inputC <- sequence[2]
	inputC <- result
	result = <-outputC

	// D
	inputD := make(chan int)
	outputD := make(chan int)
	go intcode.Execute(program, inputD, outputD)
	inputD <- sequence[3]
	inputD <- result
	result = <-outputD

	// E
	inputE := make(chan int)
	outputE := make(chan int)
	go intcode.Execute(program, inputE, outputE)
	inputE <- sequence[4]
	inputE <- result
	result = <-outputE

	if result > maxSignal {
		maxSignal = result
	}

	for {
		// A
		inputA <- result
		result = <-outputA

		// B
		inputB <- result
		result = <-outputB

		// C
		inputC <- result
		result = <-outputC

		// D
		inputD <- result
		result = <-outputD

		// E
		inputE <- result
		result = <-outputE

		if result > maxSignal {
			maxSignal = result
		}
	}
}
