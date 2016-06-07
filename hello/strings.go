package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "an implicitly typed string"
	fmt.Printf("str1 %v,%t \n", str1, str1)

	str2 := "an explicitly typed string"
	fmt.Printf("str2 %v,%t \n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	lvalue := "hello"
	uvalue := "HELLO"

	fmt.Println("Equal? ", (lvalue == uvalue))
	fmt.Println("Equal Non case sense? ", strings.EqualFold(lvalue, uvalue))
	fmt.Println("Conain exp?", strings.Contains(str1, "exp"))
	fmt.Println("Conain exp?", strings.Contains(str2, "exp"))

}
