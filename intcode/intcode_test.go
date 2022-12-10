package intcode

import (
	"testing"
)

var input chan int
var output chan int

func TestExecuteBasic(t *testing.T) {
	input = make(chan int)
	output = make(chan int)
	program := "1,9,10,3,2,3,11,0,99,30,40,50"
	expected := "3500,9,10,70,2,3,11,0,99,30,40,50"
	result, _ := Execute(program, input, output)
	if result != expected {
		t.Fatalf("\n`%s` did not match the expected result:\n`%s`", result, expected)
	}

	input = make(chan int)
	output = make(chan int)
	program = "1,0,0,0,99"
	expected = "2,0,0,0,99"
	result, _ = Execute(program, input, output)
	if result != expected {
		t.Fatalf("\n`%s` did not match the expected result:\n`%s`", result, expected)
	}

	input = make(chan int)
	output = make(chan int)
	program = "2,3,0,3,99"
	expected = "2,3,0,6,99"
	result, _ = Execute(program, input, output)
	if result != expected {
		t.Fatalf("\n`%s` did not match the expected result:\n`%s`", result, expected)
	}

	input = make(chan int)
	output = make(chan int)
	program = "2,4,4,5,99,0"
	expected = "2,4,4,5,99,9801"
	result, _ = Execute(program, input, output)
	if result != expected {
		t.Fatalf("\n`%s` did not match the expected result:\n`%s`", result, expected)
	}

	input = make(chan int)
	output = make(chan int)
	program = "1,1,1,4,99,5,6,0,99"
	expected = "30,1,1,4,2,5,6,0,99"
	result, _ = Execute(program, input, output)
	if result != expected {
		t.Fatalf("\n`%s` did not match the expected result:\n`%s`", result, expected)
	}
}

func TestExecuteInputAndOutput(t *testing.T) {
	input = make(chan int)
	output = make(chan int)
	program := "3,0,4,0,99"
	go Execute(program, input, output)
	input <- 72
	result := <-output
	if result != 72 {
		t.Fatalf("\n`%d` does not match the given input, `72`", result)
	}
}

func TestExecuteImmediateMode(t *testing.T) {
	input = make(chan int)
	output = make(chan int)
	program := "101,33,4,4,66"
	expected := "101,33,4,4,99"
	result, err := Execute(program, input, output)
	if err != nil {
		t.Fatalf("\nunexpected error `%s`", err)
	}
	if result != expected {
		t.Fatalf("\n`%s` did not match the expected result:\n`%s`", result, expected)
	}

	input = make(chan int)
	output = make(chan int)
	program = "1002,4,3,4,33"
	expected = "1002,4,3,4,99"
	result, err = Execute(program, input, output)
	if err != nil {
		t.Fatalf("\nunexpected error `%s`", err)
	}
	if result != expected {
		t.Fatalf("\n`%s` did not match the expected result:\n`%s`", result, expected)
	}
}

func TestExecuteJumpCodes(t *testing.T) {
	program := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"

	input = make(chan int)
	output = make(chan int)
	go Execute(program, input, output)
	input <- 7
	result := <-output
	if result != 999 {
		t.Fatalf("\n`%d` did not match the expected result: `999`", result)
	}

	input = make(chan int)
	output = make(chan int)
	go Execute(program, input, output)
	input <- 8
	result = <-output
	if result != 1000 {
		t.Fatalf("\n`%d` did not match the expected result: `1000`", result)
	}

	input = make(chan int)
	output = make(chan int)
	go Execute(program, input, output)
	input <- 9
	result = <-output
	if result != 1001 {
		t.Fatalf("\n`%d` did not match the expected result: `1001`", result)
	}
}
