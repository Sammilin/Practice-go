package main

import (
	"fmt"
)

func main() {
	var p *int
	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("p is nil")
	}

	var v int = 42
	p = &v

	fmt.Println("v is ", v)

	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("p is nil")
	}

	var value1 float64 = 42.23
	pointer1 := &value1
	fmt.Println("value1(*pointer) is", *pointer1)
	*pointer1 = *pointer1 / 30
	fmt.Println("value1(*pointer) is", *pointer1)
	fmt.Println("value1 is", value1)

}
