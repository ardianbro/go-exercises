// Structures
package main

import "fmt"

func main() {
	type XYZ struct {
		X int
		Y int
		Z int
	}

	var s1 XYZ // struct will be initialized with zero values
	fmt.Println(s1.X, s1.Y, s1.Z)

	p1 := XYZ{23, 12, -2}
	p2 := XYZ{Z: 12, Y: 13}
	fmt.Println(p1)
	fmt.Println(p2)

	pArr := [4]XYZ{}
	pArr[2] = p1
	pArr[0] = p2
	fmt.Println(pArr)
	p2 = XYZ{1, 2, 3}
	// changing the value of the original structure will have no effect on the objects of the array because when you
	// assign a structure to an array of structures, the structure is copied into the array
	fmt.Println(pArr)
}
