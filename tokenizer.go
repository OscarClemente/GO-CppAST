package main

import "fmt"

const letters = "abcdefghijklmnopqrstuvwxyz"
const validChars = letters + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "_0123456789$"
const hexDigits = "0123456789abcdefABCDEF"
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

func GetTokens() { //dat []byte) {
	b := []byte("ABCc€")
	if stringInSlice('c', b) {
		fmt.Println("FOUND")
	} else {
		fmt.Println("NOT FOUND")
	}
}
