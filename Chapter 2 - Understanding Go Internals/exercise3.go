// Finding out the information about the current Go environment using the functions and properties of the runtime package
// and create simple requirement check function

package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a ", runtime.GOARCH, "machine")
	fmt.Println("Using Go version ", runtime.Version())
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	fmt.Println("Number of Goroutines: ", runtime.NumGoroutine())
	requirementCheck()
}

func requirementCheck() {
	myVersion := runtime.Version()
	// the Split function on strings package will split the string of go version eg: go1.12.1 separated by ".",
	// it will create go1 12 7
	major := strings.Split(myVersion, ".")[0][2] // get the first element of array (index 0) and get the third element (index 2)
	minor := strings.Split(myVersion, ".")[1]    // get the second element of array (index 1)
	// the major variable will print the ASCII value of 1 = 049 because we slice it from a string with char in it,
	// so we need to define this as a string and convert into integer using Atoi function in strconv package
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)

	if m1 == 1 && m2 < 8 {
		fmt.Println("Need Go version 1.8 or higher!")
		return
	}

	fmt.Println("You are using Go version 1.8 or higher! ")
}
