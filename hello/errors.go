package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("filename.txt")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println("fail, error is ", err)
		fmt.Println("fail, error from error function is ", err.Error())
	}

	my_error := errors.New("This is custom err")
	fmt.Println(my_error)

	attendance := map[string]bool{
		"Mike": true,
		"Amy":  false}

	attended, ok := attendance["M"] //ok=true means there's no error

	if ok {
		fmt.Println("Mike is attend ? ", attended)
	} else {
		fmt.Println("Mike not inside")
	}
}
