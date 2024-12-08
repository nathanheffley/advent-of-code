package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/intcode"
)

func main() {
	program := "203,80,109,1,204,79,99"

	input := make(chan int)
	output := make(chan int)
	go intcode.Execute(program, input, output)
	input <- 66

	for msg := range output {
		fmt.Println(msg)
	}

}
