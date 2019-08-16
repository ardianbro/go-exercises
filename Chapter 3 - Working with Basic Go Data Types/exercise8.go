// Go constants, variables that cannot change their values
package main

import "fmt"

// type is a way of defining a new named type that use the same underlying type as existing one.
// this is mainly used for differentiating between different types that might use the same kind of data.
type Digit int
type Power2 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	// using iota in constants declaration
	const (
		Zero Digit = iota
		One
		Two
		Three
		Four

		// The use of iota is equivalent to following declaration of constants below:
		// Zero = 0
		// One = 1
		// Two = 2
		// Three = 3
		// Four = 4
	)
	fmt.Println(One)
	fmt.Println(Two)

	const (
		p2_0 Power2 = 1 << iota // bitwise operation left shifting (1 << 0)
		_                       // again, _ (underscore) is used to ignore thing
		p2_2                    // (1 << 2) = 100 (binary) = 4
		_                       // ignored
		p2_4                    // (1 << 4) = 10000 (binary) = 16
		_                       // ignored
		p2_6                    // (1 << 6) = 1000000 (binary) = 64
	)
	fmt.Println("2^0: ", p2_0)
	fmt.Println("2^2: ", p2_2)
	fmt.Println("2^4: ", p2_4)
	fmt.Println("2^6: ", p2_6)
}
