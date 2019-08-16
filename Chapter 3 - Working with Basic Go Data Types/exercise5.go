// Understanding more of Go slices

package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)
	integer := make([]int, 2) // the slice will be initiated automatically, so it will be [0 0]
	fmt.Println(integer)
	integer = nil // set the slice to empty slice
	fmt.Println(integer)
	fmt.Println()

	anArray := [5]int{-1, -2, -3, -4, -5}
	refAnArray := anArray[:]
	fmt.Println(anArray)
	fmt.Println(refAnArray)
	anArray[4] = -50        // set the value of the fifth element in array anArray to -5
	fmt.Println(refAnArray) // the refAnArray value will also changed because its referencing to the same memory address
	fmt.Println()

	s := make([]byte, 5)
	fmt.Println(s) // the slice automatically initialized by go ([0 0 0 0 0])
	twoD := make([][]int, 3)
	fmt.Println(twoD) // the default zero value of the slices is nil, the elements of multi-dimensional slice are slices
	fmt.Println()

	// initialize all the elements of a slice with two dimensions manually
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 4; j++ { // define length of each slice element on condition expression
			twoD[i] = append(twoD[i], i*j)
		}
	}
	fmt.Println(twoD) // check the initialized two dimensional slice

	// print all of the elements of a slice with two dimension
	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i: ", i, "val: ", y)
		}
		fmt.Println()
	}
}
