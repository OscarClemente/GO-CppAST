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
	fmt.Println(datStr)
	val := datStr[2]
	if string(val) == "AS" {
		fmt.Println("hmmm")
	} else {
		fmt.Println("failed", val)
	}
	fmt.Println("--------------------")
	tokenSlice := GetTokens(datStr)
	//fmt.Println(tokenSlice)
	for _, token := range tokenSlice {
		fmt.Println(token)
	}
}
