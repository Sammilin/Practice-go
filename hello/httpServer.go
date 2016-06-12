package main

import (
	"fmt"
	"net/http"
)

type Hello struct {
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>This is HTTP Server, Produce by Go</h1>")
}

func main() {
	var h Hello

	err := http.ListenAndServe("localhost:4000", h)
	CheckErr(err)

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
