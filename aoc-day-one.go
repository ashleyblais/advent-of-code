package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func runPartOne(values []string) {
	var sum = 0

	for i := 0; i < len(values); i++ {
		value, conversionErr := strconv.ParseFloat(values[i], 32)

		if conversionErr != nil {
			fmt.Print(conversionErr)
		} else {
			sum += (int(value/3) - 2)
		}
	}

	fmt.Println(sum)
}

func runPartTwo(values []string) {
	var sum = 0

	for i := 0; i < len(values); i++ {
		value, conversionErr := strconv.ParseFloat(values[i], 32)

		if conversionErr != nil {
			fmt.Print(conversionErr)
		} else {
			var fuelNeeded = (int(value/3) - 2)
			sum += fuelNeeded

			for fuelNeeded > 0 {
				fuelNeeded = (int(fuelNeeded/3) - 2)

				if fuelNeeded > 0 {
					sum += fuelNeeded
				} else {
					fuelNeeded = 0
				}
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	data, fileReadErr := ioutil.ReadFile("./day-one-input.txt")

	if fileReadErr != nil {
		fmt.Print(fileReadErr)
	}

	values := strings.Split(string(data), "\n")

	runPartOne(values)
	runPartTwo(values)
}
