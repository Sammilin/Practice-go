package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")

	str1 := "the test1 "
	str2 := "print test2 xxxx"
	str3 := " test3333333"
	stringLength, _ := fmt.Println(str1, str2, str3)

	aNumber := 34
	isTrue := true

	// if err == nil {
	fmt.Println("string len:", stringLength)
	// }

	fmt.Printf("this is number %v\n", aNumber)

	fmt.Printf("this is boolean true %v\n", isTrue)

	fmt.Printf("this is float %.2f\n", float64(aNumber))

	fmt.Printf("this is data type %T,%T,%T,%T,%T\n", str1, str2, str3, aNumber, isTrue)

	myDatatype := fmt.Sprintf("this is data type %T,%T,%T,%T,%T", str1, str2, str3, aNumber, isTrue)
	fmt.Println(myDatatype)

}
