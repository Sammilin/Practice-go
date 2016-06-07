package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var s string			//explicitly declaration
	// fmt.Scanln(&s)
	// fmt.Println(s)

	const anInteger int = 2           //explicit
	const asString = "this is string" //implicit

	reader := bufio.NewReader(os.Stdin) //implicity type declasration
	fmt.Println("Please type text:")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Println("Please type Number:")
	str, _ = reader.ReadString('\n') //str has alreday declare, didn't need to := again.
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("value of number: ", f)
	}

}
