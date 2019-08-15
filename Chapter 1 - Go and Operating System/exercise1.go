// go program that finds the sum of all of its numeric command-line arguments
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	var total float64

	for i := 1; i < len(arguments); i++ {
		arg, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			total = total + arg
		}
	}

	fmt.Println(total)
}
