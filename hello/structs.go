package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	weight int
}

func main() {
	poodle := Dog{"poddle", 34}

	fmt.Printf("%+v \n", poodle)

	fmt.Printf("breed: %v, weight: %v\n", poodle.Breed, poodle.weight)

}
