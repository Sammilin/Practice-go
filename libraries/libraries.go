package main

import (
	"fmt"
	"stringutil"
)

func main() {
	//multi return
	f1, n1 := stringutil.Fullname("Sammi", "Lin")
	fmt.Printf("fullname is %v, char length is %v \n", f1, n1)
	f2, n2 := stringutil.FullnameNakedResult("PeiShan", "Lin")
	fmt.Printf("fullname is %v, char length is %v \n", f2, n2)
}
