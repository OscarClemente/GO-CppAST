package main

import (
	"fmt"
)

const intOrFloatDigits = "01234567890eE-+"

var strPrefixes = []string{"R", "u8", "u8R", "u", "uR", "U", "UR", "L", "LR"}

func stringInSlice(a byte, list []byte) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetTokens(source []byte) {
	letters := []byte("abcdefghijklmnopqrstuvwxyz")
	lettersUpper := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numChars := []byte("_0123456789$")
	validChars := append(letters[:], lettersUpper[:]...)
	validChars = append(validChars, numChars...)
	hexDigits := []byte("0123456789abcdefABCDEF")

	fmt.Println(hexDigits)

	i := 0
	end := len(source)

	for i < end {
		// skip spaces
		for i < end && source[i] == ' ' {
			i++
		}
		if i >= end {
			return
		}

		/*tokenType := Unknown
		start := i
		c := dat[i]

		if unicode.IsLetter(c) {

		}*/
	}

	b := []byte("ABC€")
	if stringInSlice('c', b) {
		fmt.Println("FOUND")
	} else {
		fmt.Println("NOT FOUND")
	}
}
