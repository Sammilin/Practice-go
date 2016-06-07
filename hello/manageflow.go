package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// var x float64 = 42
	var result string

	if x := 10; x > 10 {
		result = "x bigger than 10"
	} else if x == 10 {
		result = "x equal 10"
	} else {
		result = "x samller than 10"
	}

	fmt.Println(result)

	rand.Seed(time.Now().Unix())
	// dow := rand.Intn(6) + 1
	// fmt.Println(dow)

	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "Sunday"
	case 7:
		result = "Staturday"
	default:
		result = "This is week day"
	}
	fmt.Println(result)

	y := 42

	switch {
	case y > 0:
		result = "y > 0"
		// fallthrough
	case y == 0:
		result = "y equal 0"
	default:
		result = "normal"

	}
	fmt.Println(result)

	//Loop
	colors := []string{"red", "yellow", "blue"}
	for i := range colors {
		fmt.Println(colors[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("sum =", sum)
		if sum > 1000 {
			goto labelgotosample
		}
		if sum > 500 {
			break
		}

	}

labelgotosample:
	fmt.Println("end of program")

}
