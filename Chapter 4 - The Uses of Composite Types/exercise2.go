// Pointer to Sructures
package main

import "fmt"

type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

// creating a new structure using this approach has many advantages over initializing variables manually, allowing us to
// check wether information provided is both correct and valid
func newStructurePointer(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}
	return &myStructure{n, s, h}
}

// no pointer version of NewStruct()
func newStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}
	return myStructure{n, s, h}
}

func main() {
	s1 := newStructurePointer("Ardian", "Pradipta", 178)
	s2 := newStructure("Ardian", "Pradipta", 178)
	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)

	pS := new(myStructure)
	// create a fresh myStructure variable, using new keyword will return the memory address of the allocated object
	fmt.Println(pS)
}
