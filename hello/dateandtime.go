package main

import (
	"fmt"
	"time"
)

func main() {
	// time.UTC()
	t := time.Now().UTC()
	fmt.Println("This is UTC time " + t.String())

	fmt.Printf("this month is %v\n", t.Month())

	longFormat := "Monday, January 2,2006"
	shortFormat := "1/1/06"
	fmt.Printf("logformat is %v\n", t.Format(longFormat))
	fmt.Println(t.Format(shortFormat))
}
