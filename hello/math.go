package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	f1, f2, f3 := 10.3, 11.2, 12.2
	fsum := f1 + f2 + f3
	fmt.Println(fsum)

	var b1, b2, b3, bsum big.Float
	b1.SetFloat64(12.4)
	b2.SetFloat64(14.1)
	b3.SetFloat64(10.2)

	bsum.Add(&b1, &b2).Add(&bsum, &b3)
	fmt.Printf("BigSum = %.10g\n", &bsum)

	mradius := 15.5
	circumference := mradius * math.Pi
	fmt.Printf("circum ference is %.2f\n", circumference)
}
