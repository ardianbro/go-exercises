// The switch statement

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("usage: ", filepath.Base(arguments[0]), "number")
		os.Exit(1)
	}

	number, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("This value not an integer:", number)
	} else {
		switch {
		case number < 0:
			fmt.Println("Less than zero!")
		case number > 0:
			fmt.Println("Bigger than zero!")
		default:
			fmt.Println("Zero")
		}
	}

	asString := arguments[1]
	switch asString {
	case "5":
		fmt.Println("Five!")
	case "0":
		fmt.Println("Zero!")
	default:
		fmt.Println("Do not care!")
	}

	var negative = regexp.MustCompile("-")
	var floatingPoint = regexp.MustCompile(`\d?\.\d`)
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`) // any char execpt @, then @ followed by any char except @, then . followed by any char except @

	switch {
	case negative.MatchString(asString):
		fmt.Println("Negative number")
	case floatingPoint.MatchString(asString):
		fmt.Println("Floating point")
	case email.MatchString(asString):
		fmt.Println("It is an email")
		fallthrough // tells go to execute the branch that follows the current one, in this case the default branch
	default:
		fmt.Println("Something else")
	}

	var aType error = nil
	switch aType.(type) {
	case nil:
		fmt.Println("It is nil interface!")
	default:
		fmt.Println("Not nil interface!")
	}

}
