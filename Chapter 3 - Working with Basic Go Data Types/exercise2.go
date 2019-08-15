// Multi-dimensional arrays in Golang

package main

import "fmt"

func main() {
	oneD := [4]int{1, 2, 3, 4}
	twoD := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	threeD := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	fmt.Println("The length of", oneD, "is", len(oneD))
	fmt.Println("The length of", twoD, "is", len(twoD))
	fmt.Println("The length of", threeD, "is", len(threeD))

	// using basic for loop
	for i := 0; i < len(threeD); i++ {
		val1 := threeD[i]
		for j := 0; j < len(val1); j++ {
			val2 := val1[j]
			for k := 0; k < len(val2); k++ {
				fmt.Print(val2[k], " ")
			}
		}
		println()
	}

	// using for loop with range to iterate over an array, _ is used to ignore something, in this case is the index of the iterated array
	for _, val1 := range threeD {
		for _, val2 := range val1 {
			for _, val3 := range val2 {
				fmt.Print(val3, " ")
			}
		}
		println()
	}
}
