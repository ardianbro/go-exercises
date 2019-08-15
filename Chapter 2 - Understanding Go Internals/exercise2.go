// Understanding panic and recover functions

package main

import "fmt"

func a() {
	fmt.Println("This is inside function a()")
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Recovering inside of function a()!")
		}
	}()
	fmt.Println("About to call function b()")
	b()
	// Two line of code below will be unreachable because the b function is panicked thus make the function a() finished
	// early, and then leaves the deferred anonymous function that calls recover() ran
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("This is inside function b()")
	panic("Panic in b!")
	// The code below is unreachable because the current program flow terminated immediately after panic() have been called
	fmt.Println("Exiting b()")
}

func main() {
	a()
	fmt.Println("main() has ended!")
}
