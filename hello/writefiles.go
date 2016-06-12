package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go"

	file, err := os.Create("./fromstring.txt")
	CheckErr(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	CheckErr(err)
	fmt.Printf("all file is write done, len is %v", ln)

	bytes := []byte(content)
	ioutil.WriteFile("bytefiles.txt", bytes, 0644)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
