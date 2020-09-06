package main

import (
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
	/*fmt.Print(dat)
	fmt.Println()
	for i, c := range dat {
		if c == '\n' {
			fmt.Printf("%d, \\n\n", i)
		} else {
			fmt.Printf("%d, %c\n", i, c)
		}
	}
	fmt.Println()
	fmt.Print(string(dat))
	fmt.Println()
	datstring := string(dat)
	for i, c := range datstring {
		if c == '\n' {
			fmt.Printf("%d, \\n\n", i)
		} else {
			fmt.Printf("%d, %c\n", i, c)
		}
	}
	fmt.Printf("%c <-\n", datstring[50])
	fmt.Println(len(datstring))
	fmt.Println(utf8.RuneCountInString(datstring))

	s2 := token{}
	fmt.Println(s2)
	s2.print()

	fmt.Println("--------")
	datrunes := []rune(datstring)
	for i, c := range datrunes {
		if c == '\n' {
			fmt.Printf("%d, \\n\n", i)
		} else {
			fmt.Printf("%d, %c\n", i, c)
		}
	}
	fmt.Printf("%c <-\n", datrunes[49])
	fmt.Println(len(datrunes))

	comilla := []byte("'")
	if comilla[0] == '\'' {
		fmt.Println("YEEES")
	}*/
	GetTokens(dat)
}
