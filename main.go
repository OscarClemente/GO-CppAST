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
	fmt.Print(string(dat))
	fmt.Println()
	/*for i, c := range dat {
		if c == '\n' {
			fmt.Printf("%d, \\n\n", i)
		} else {
			fmt.Printf("%d, %c\n", i, c)
		}
	}*/
	token()
}
