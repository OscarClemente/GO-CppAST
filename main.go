package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("files/foo.hpp")
	check(err)
	datStr := string(dat)

	tokenSlice := GetTokens(datStr)
	for _, token := range tokenSlice {
		fmt.Println(token)
	}
}
