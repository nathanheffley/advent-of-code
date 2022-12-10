package intcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Execute(program string, input chan int, output chan int) (string, error) {
	strInts := strings.Split(program, ",")
	ints := make([]int, len(strInts))
	for i, strInt := range strInts {
		ints[i], _ = strconv.Atoi(strInt)
	}

	i := 0
	for {
		commandInt := ints[i]

		if commandInt == 99 {
			close(input)
			close(output)
			break
		}

		fullCommand := fmt.Sprintf("%05d", commandInt)

		command := fullCommand[3:]

		arg1Mode := fullCommand[2:3]
		rawArg1 := func() int {
			return ints[i+1]
		}
		arg1 := func() int {
			if arg1Mode == "0" {
				return ints[rawArg1()]
			} else if arg1Mode == "1" {
				return rawArg1()
			}
			return -math.MinInt32
		}

		arg2Mode := fullCommand[1:2]
		rawArg2 := func() int {
			return ints[i+2]
		}
		arg2 := func() int {
			if arg2Mode == "0" {
				return ints[rawArg2()]
			} else if arg2Mode == "1" {
				return rawArg2()
			}
			return -math.MinInt32
		}

		// arg3Mode := fullCommand[:1]
		rawArg3 := func() int {
			return ints[i+3]
		}
		// arg3 := func() int {
		// 	if arg3Mode == "0" {
		// 		return ints[rawArg3()]
		// 	} else if arg3Mode == "1" {
		// 		return rawArg3()
		// 	}
		// 	return -math.MinInt32
		// }

		if command == "01" {
			ints[rawArg3()] = arg1() + arg2()
			i += 4
			continue
		}

		if command == "02" {
			ints[rawArg3()] = arg1() * arg2()
			i += 4
			continue
		}

		if command == "03" {
			in := <-input
			ints[rawArg1()] = in
			i += 2
			continue
		}

		if command == "04" {
			output <- arg1()
			i += 2
			continue
		}

		if command == "05" {
			if arg1() != 0 {
				i = arg2()
				continue
			}
			i += 3
			continue
		}

		if command == "06" {
			if arg1() == 0 {
				i = arg2()
				continue
			}
			i += 3
			continue
		}

		if command == "07" {
			if arg1() < arg2() {
				ints[rawArg3()] = 1
			} else {
				ints[rawArg3()] = 0
			}
			i += 4
			continue
		}

		if command == "08" {
			if arg1() == arg2() {
				ints[rawArg3()] = 1
			} else {
				ints[rawArg3()] = 0
			}
			i += 4
			continue
		}

		fmt.Println(ints)
		fmt.Println(fullCommand)
		return "", fmt.Errorf("an invalid command was received")
	}

	for i, int := range ints {
		strInts[i] = fmt.Sprintf("%d", int)
	}
	return strings.Join(strInts, ","), nil
}
