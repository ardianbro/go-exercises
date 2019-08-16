// Sorting slices using sort.Slice()

package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStructure, 0)
	// Adding elements into the slice
	mySlice = append(mySlice, aStructure{"Ardian", 175, 75})
	mySlice = append(mySlice, aStructure{"Hasta", 178, 76})
	mySlice = append(mySlice, aStructure{"Singgih", 168, 60})
	mySlice = append(mySlice, aStructure{"Fatah", 180, 55})
	fmt.Println("0: ", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		// sort by height from short to tall
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		// sort by height from tall to short
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)

}
