// The strings package

package main

import (
	"fmt"
	s "strings" // give alias to strings package
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))

	f("%s\n", s.Title(s.ToLower("tHis wILL be A TitlE!")))

	f("EqualFold: %v\n", s.EqualFold("Ardian", "ARDIan"))
	f("EqualFold: %v\n", s.EqualFold("Ardian", "ardia"))

	f("Prefix: %v\n", s.HasPrefix("Ardian", "Ar")) // check whether the string starts with the substring
	f("Prefix: %v\n", s.HasPrefix("Ardian", "ar"))
	f("Suffix: %v\n", s.HasSuffix("Ardian", "di")) // check whether the string ends with the substring
	f("Suffix: %v\n", s.HasSuffix("Ardian", "di"))
	f("Index: %v\n", s.Index("Ardian", "an")) //Index returns the index of the first instance of substr in s, or -1 if substr is not present in s
	f("Index: %v\n", s.Index("Ardian", "An"))
	f("Count: %v\n", s.Count("Ardian", "a")) // Count counts the number of non-overlapping instances of substr in s
	f("Count: %v\n", s.Count("Ardian", "A"))
	f("Repeat: %s\n", s.Repeat("ab", 5))

	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s", s.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t "))

	// strings.Compare() compares two strings lexicographically, return 0 if identical, and -1 or 1 otherwise
	f("Compare: %v\n", s.Compare("Ardian", "ARDIAN"))
	f("Compare: %v\n", s.Compare("Ardian", "Ardian"))
	f("Compare: %v\n", s.Compare("ARDIAN", "ARDIan"))
	f("Fields: %v\n", s.Fields("This is a string!")) //split string with whitespace as separator
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))
	f("%s\n", s.Split("abcd efg", "")) // split each character

	f("%s\n", s.Replace("abcd efg", "", "_", -1)) // negative value on 4th parameter means there is no limit to the number of replacement
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))                   // Join concatenates the elements of a to create a single string
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++")) // SplitAfter slices s into all substrings after each instance of sep

	trimFunction := func(c rune) bool {
		println(string(c), !unicode.IsLetter(c))
		return !unicode.IsLetter(c) // will return all unicode char that is not letter to be trimmed
	}

	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction)) // function splits its first parameter string into substrings based on the separator string that is given as the second parameter to the function
}
