// Dealing with times and dates

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Epoch time:", time.Now().Unix()) // will return Unix epoch time, the number of seconds elapsed since 00:00:00 UTC, 1 Jan 1970
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339)) // The Format() function allows conversion of a time variable to another format
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second) // easy way to emulating a delay from the execution of a true function
	t1 := time.Now()
	fmt.Println("Time difference:", t1.Sub(t)) // Sub() function used to find the time difference between two times

	formatT := t.Format("01 January 2006") // formatting t variable to match the layout given
	fmt.Println(formatT)
	loc, _ := time.LoadLocation("Asia/Jakarta")
	yogyakartaTime := t.In(loc)
	fmt.Println("Yogyakarta:", yogyakartaTime)
}
