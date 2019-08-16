// Understanding more about Go pointers
package main

import "fmt"

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int { // return a pointer
	v := n * n
	return &v
}

func main() {
	i := -10
	j := 25

	pI := &i // pI pointing to memory address of i
	pJ := &j // pJ pointing to memory address of j

	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)

	*pI = 123456
	*pI--
	fmt.Println("i:", i)

	getPointer(pJ)         // any changes made to pJ variable inside getPointer() will have an effect on the value of j, because pJ points to j
	fmt.Println("j:", j)   // verify the changes
	k := returnPointer(12) // returns a pointer that is assigned to the k pointer variable
	fmt.Println(*k)
	fmt.Println(k)
}
