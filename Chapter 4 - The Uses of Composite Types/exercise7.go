// Rune, an int32 value, a Go type that used for representing a Unicode code point

package main

import "fmt"

func main() {
	const r1 = '€'
	fmt.Println("(int32) r1:", r1)
	fmt.Printf("(HEX) r1: %x\n", r1)
	fmt.Printf("(as a String) r1: %s\n", r1)
	fmt.Printf("(as a Character) r1: %c\n", r1)

	fmt.Println("A string is a collection of runes:", []byte("Ardian"))
	aString := []byte("Ardian")
	for x, y := range aString {
		fmt.Println(x, y)
		fmt.Printf("Char: %c\n", aString[x])
	}

	fmt.Printf("%s\n", aString)
}
