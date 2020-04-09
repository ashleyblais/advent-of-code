package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	partOne := runPartOne(12, 2)
	fmt.Println("PART ONE ANSWER: ", partOne[0])

	runPartTwo()
}

func runPartOne(inputOne, inputTwo int64) []int64 {
	values := getComputerData()

	values[1] = inputOne
	values[2] = inputTwo

	for i := 0; i < len(values); i += 4 {
		var opcode = values[i]

		if opcode != 1 && opcode != 2 && opcode != 99 {
			fmt.Println()
		}

		var arraySection []int64

		if opcode == 99 {
			break
		} else {
			arraySection = values[i:(i + 4)]
			var valueOne = values[arraySection[1]]
			var valueTwo = values[arraySection[2]]

			if opcode == 1 {
				output := valueOne + valueTwo
				values[arraySection[3]] = output
			}

			if opcode == 2 {
				output := (valueOne * valueTwo)
				values[arraySection[3]] = output
			}
		}
	}

	return values
}

func runPartTwo() {
	values := getComputerData()

	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values); j++ {
			values = runPartOne(int64(j), int64(i))

			noun := values[1]
			verb := values[2]

			if values[0] == 19690720 {
				fmt.Println("PART TWO ANSWER: ", 100*noun+verb, "INPUTS: ", noun, verb)
			}
		}
	}
}

func getComputerData() []int64 {
	data, fileReadErr := ioutil.ReadFile("./day-two-input.txt")

	if fileReadErr != nil {
		fmt.Print(fileReadErr)
	}

	values := strings.Split(string(data), ",")
	var intValues []int64

	for i := 0; i < len(values); i++ {
		var convertedInt, conversionErr = strconv.ParseInt(values[i], 0, 64)
		if conversionErr != nil {
			fmt.Println(conversionErr)
		} else {
			intValues = append(intValues, convertedInt)
		}
	}

	return intValues
}
