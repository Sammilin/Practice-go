package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "wan"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "meow"
}

func main() {
	poodle := Dog{}
	fmt.Println(poodle.Speak())

	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

}
