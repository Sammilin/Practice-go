package main

import (
	"fmt"
)

func main() {
	doSomething()

	fmt.Println(addValue(12, 34))

	fmt.Println(addValues(12, 22, 22))

	//multi return
	f1, n1 := fullname("Sammi", "Lin")
	fmt.Printf("fullname is %v, char length is %v \n", f1, n1)
	f2, n2 := fullnameNakedResult("PeiShan", "Lin")
	fmt.Printf("fullname is %v, char length is %v \n", f2, n2)
}

func doSomething() {
	fmt.Println("get success")
}

func addValue(val1, val2 int) int {
	sum := 0
	sum = val1 + val2
	return sum
}

func addValues(values ...int) int {
	fmt.Printf("format %T\n", values)

	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

func fullname(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func fullnameNakedResult(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
