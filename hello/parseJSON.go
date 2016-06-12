package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Tour struct {
	Name, Price string
}

func main() {
	url := "http://services.explorecalifornia.org/json/tours.php"

	content := readRespContent(url)
	tours := toursFromJson(content)

	for i, tour := range tours {
		fmt.Println(i, " name : ", tour.Name, " price: ", tour.Price)
	}

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readRespContent(url string) string {

	resp, err := http.Get(url)
	CheckErr(err)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)

	return string(bytes)
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20) //make slice of tours, 20 just guess, no reason.

	decoder := json.NewDecoder(strings.NewReader(content)) //decoder the string for JSON function
	_, err := decoder.Token()
	CheckErr(err)

	var tour Tour
	for decoder.More() { //like object.hasNext()?
		err := decoder.Decode(&tour) //&tour -> it's a reference, decode structure and refer the value
		CheckErr(err)
		tours = append(tours, tour)
	}

	return tours

}
