package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func runPartOne(values []int64) {
	for i := 0; i < len(values); i += 4 {
		var opcode = values[i]

		if opcode != 1 && opcode != 2 && opcode != 99 {
			fmt.Println()
		}

		var arraySection []int64

		if opcode == 99 {
			fmt.Println("PROGRAM ENDED")
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

	fmt.Println(values[0])
}

func main() {
	data, fileReadErr := ioutil.ReadFile("./day-two-input.txt")

	if fileReadErr != nil {
		fmt.Print(fileReadErr)
	}

	values := strings.Split(string(data), ",")

	var intValues []int64

	// converting the values from the text files to ints for easier usage
	for i := 0; i < len(values); i++ {
		var convertedInt, conversionErr = strconv.ParseInt(values[i], 0, 64)
		if conversionErr != nil {
			fmt.Println(conversionErr)
		} else {
			intValues = append(intValues, convertedInt)
		}
	}

	intValues[1] = 12
	intValues[2] = 2

	runPartOne(intValues)
}
