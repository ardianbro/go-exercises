// Understanding defer keyword

package main

import "fmt"

func defer1() {
	for i := 5; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func defer2() {
	for i := 5; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}

func defer3() {
	for i := 5; i > 0; i-- {
		defer func(i int) {
			fmt.Print(i, " ")
		}(i)
	}
}

func main() {
	defer1() // will return 1 2 3 4 5 as output, because of the LIFO order of defer, so the first number will be printed is 5 and then 4 and so on
	fmt.Println()
	defer2() // will return 0 0 0 0 0 as output, because the value that passed into the anonymous function is the value of last iteration, that is 0
	fmt.Println()
	defer3() // will return 1 2 3 4 5 as output, because the value will be passed into anonymous function on each iteration, only the print function is called last with LIFO order
}
