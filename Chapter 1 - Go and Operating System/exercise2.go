// go program that finds the average value of all of its float command-line arguments
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
	var average float64
	var argcount float64

	for i := 1; i < len(arguments); i++ {
		arg, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			total = total + arg
			argcount++
		}
	}

	average = total / argcount
	fmt.Println(total)
	fmt.Println(argcount)
	fmt.Println(average)
}
