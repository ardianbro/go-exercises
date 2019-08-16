// for loops in Golang

package main

import (
	"fmt"
)

func main() {
	loop1()
	println()
	loop2()
	println()
	loop3()
	println()
	loop4()

}

// simples form of loop
// init statement: executed before first iteration; condition expression: evaluated before every iteration; post statement: executed at the end of every iteration
func loop1() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue // the loop will be skipped on i divisible by 20 eg: 20, 40, etc, so print function will not be called
		}

		if i == 95 {
			break // the loop will be stopped on i == 95
		}

		fmt.Print(i, " ")
	}
}

// while (condition) loop emulation / check the condition first then do the work()
func loop2() {
	i := 10
	for {
		if i < 0 { // check the condition
			break
		}
		fmt.Print(i, " ") // do the work()
		i--
	}
}

// do ... while loop emulation / do the work() first then check the condition
func loop3() {
	i := 0
	anExpression := true
	for ok := true; ok; ok = anExpression { // checking the condition
		if i > 10 {
			anExpression = false // change the condition (the loop still working after unfulfilled condition created)
		}

		fmt.Print(i, " ") // do the work
		i++
	}
}

// for loop using range to iterate over an array
func loop4() {
	anArray := [5]int{0, 1, -1, 2, -2} // declare an array that store 5 integers and initiate it
	for i, val := range anArray {
		fmt.Println("index:", i, " value:", val)
	}
}
