// Regular expression and pattern matching on parsing IP address

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)

}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Usage:", filepath.Base(arguments[0]), "logFile")
		os.Exit(1)
	}

	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file", err)
			os.Exit(-1)
		}
		defer func() {
			err := f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
		}()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("Error reading file", err)
				break
			}
			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			} else {
				fmt.Println(ip)
			}

		}
	}

	// usage example: go run exercise5.go auth.log
	// you can also add traditional unix command line utilities to process the preceding output
	// usage example: go run exercise5.go auth.log | sort | uniq -c | sort
}
