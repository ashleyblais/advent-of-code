package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, fileReadErr := ioutil.ReadFile("./numbers.txt")

	if fileReadErr != nil {
		fmt.Print(fileReadErr)
	}

	values := strings.Split(string(data), "\n")

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
