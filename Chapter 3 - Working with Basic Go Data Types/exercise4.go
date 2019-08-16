// The copy() function on slices and arrays

package main

import "fmt"

func main() {
	a6 := []int{1, 2, 3, 4, 5, 6}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6: ", a6)
	fmt.Println("a4: ", a4)
	copy(a6, a4)
	fmt.Println("a6: ", a6)
	fmt.Println("a4: ", a4)
	fmt.Println()

	b6 := []int{1, 2, 3, 4, 5, 6}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6: ", b6)
	fmt.Println("b4: ", b4)
	copy(b4, b6)
	fmt.Println("b6: ", b6)
	fmt.Println("b4: ", b4)
	fmt.Println()

	array4 := [4]int{1, 2, 3, 4}
	slice6 := []int{-1, -2, -3, -4, -5, -6}
	// the array will be converted to a slice with the help of the [:] notation (array4[0:])
	copy(slice6, array4[0:])
	fmt.Println("array4: ", array4[0:])
	fmt.Println("slice6: ", slice6)
	fmt.Println()

	array5 := [5]int{-1, -2, -3, -4, -5}
	slice7 := []int{1, 2, 3, 4, 5, 6, 7}
	// again, the array will be converted to a slice with the help of the [:] notation (array5[0:])
	copy(array5[0:], slice7)
	fmt.Println("array5: ", array5)
	fmt.Println("slice7: ", slice7)
}
