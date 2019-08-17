// Regular expression and pattern matching example on log files of an Apache web server

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide one text file to process!")
		os.Exit(1)
	}

	filename := arguments[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file ", err)
		os.Exit(1)
	}
	// defer function will be called after main program is executed and with the order of LIFO
	defer func() {
		err := f.Close()
		// handling error on Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	notAMatch := 0 // Initiate variable to count line that doesn't match
	r := bufio.NewReader(f)
	// Read line by line
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading file ", err)
		}

		r1 := regexp.MustCompile(`.*\[(\d\d/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)].*`)
		if r1.MatchString(line) {
			match := r1.FindStringSubmatch(line)
			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Println(strings.Replace(line, match[1], newFormat, 1))
			} else {
				//fmt.Println(err)
				notAMatch++
			}
			continue
		}
		r2 := regexp.MustCompile(`.*\[(\w+-\d\d-\d\d:\d\d:\d\d:\d\d.*)].*`)
		if r2.MatchString(line) {
			match := r2.FindStringSubmatch(line)
			d1, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Println(strings.Replace(line, match[1], newFormat, 1))
			} else {
				//fmt.Println(err)
				notAMatch++
			}
			continue
		}
	}
	fmt.Println(notAMatch, " lines did not match")

	// usage example: go run exercise4.go LogEntries.txt
}
