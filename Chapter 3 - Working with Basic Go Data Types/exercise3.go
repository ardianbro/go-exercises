// Understanding length and capacity of slices

package main

import "fmt"

// function to help printing the slice
func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func main() {
	// create a slice and initiate it
	aSlice := []int{1, 0, -4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Capacity: %d, Length: %d \n", cap(aSlice), len(aSlice))

	// add a value to the slice
	aSlice = append(aSlice, 42)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Capacity: %d, Length: %d \n", cap(aSlice), len(aSlice))

	// add more value to the slice
	aSlice = append(aSlice, 27)
	aSlice = append(aSlice, -18)
	aSlice = append(aSlice, 67)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Capacity: %d, Lenght %d \n", cap(aSlice), len(aSlice))

	// The length of slice is growing according the number of the elements inside the slice
	// The capacity will be doubled each time we add new element on a slice when the slice has same length and capacity
}
