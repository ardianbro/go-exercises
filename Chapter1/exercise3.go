// go program that keeps reading integers until it gets the word "STOP" as input
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var f *os.File
	f = os.Stdin
	// defer function will be called after main program is executed and with the order of LIFO
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if scanner.Text() == "STOP" {
			os.Exit(0)
		} else if err != nil {
			fmt.Println(scanner.Text(), "is not a number, please input a number")
		} else {
			fmt.Println(">", x)
		}
	}
}
