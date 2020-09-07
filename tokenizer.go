package main

import (
	"bytes"
	"fmt"
)

const intOrFloatDigits = "01234567890eE-+"

var strPrefixes = []string{"R", "u8", "u8R", "u", "uR", "U", "UR", "L", "LR"}

func byteInSlice(a byte, list []byte) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getString(source []byte, start int, i int) int {
	i = bytes.Index(source[i+1:], []byte("\""))
	for source[i-1] == '\\' {
		// count trailing backslashes
		backslashCount := 1
		j := i - 2
		for source[j] == '\\' {
			backslashCount++
			j--
		}
		// when trailing backslashes are even they escape each other
		if backslashCount%2 == 0 {
			break
		}
		i = bytes.Index(source[i+1:], []byte("\""))
	}
	return i + 1
}

func getChar(source []byte, start int, i int) int {
	i = bytes.Index(source[i+1:], []byte("'"))
	for source[i-1] == '\\' {
		// Need special case '\\'
		if (i-2) > start && source[i-2] == '\\' {
			break
		}
		i = bytes.Index(source[i+1:], []byte("'"))
	}
	// Unterminated single quotes
	if i < 0 {
		i = start
	}
	return i + 1
}

func GetTokens(source []byte) []*token {
	letters := []byte("abcdefghijklmnopqrstuvwxyz")
	lettersUpper := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numChars := []byte("0123456789")
	extraChar := []byte("_$")
	alphaNumeric := append(letters, lettersUpper...)
	alphaNumeric = append(alphaNumeric, numChars...)
	validChars := append(alphaNumeric, extraChar...)
	hexDigits := []byte("0123456789abcdefABCDEF")
	intOrFloatDigits := []byte("0123456789eE-+")
	intOrFloatDigits2 := []byte("0123456789eE-+.")

	fmt.Println(hexDigits)
	ignoreErrors := false
	countIfs := 0

	i := 0
	end := len(source)
	var tokenSlice []*token

	for i < end {
		// skip spaces
		for i < end && source[i] == ' ' {
			i++
		}
		if source[i] == '\n' {
			i++
			continue
		}
		if i >= end {
			return tokenSlice
		}

		tokenType := Unknown
		start := i
		c := source[i]

		if byteInSlice(c, alphaNumeric) || c == '_' {
			tokenType = Name
			for byteInSlice(source[i], validChars) {
				i++
			}
			if source[i] == '\'' && (i-start) == 1 &&
				(source[start] == 'u' || source[start] == 'U' || source[start] == 'L') {
				fmt.Println(tokenType, start) // delete
				tokenType = Constant
				i = getChar(source, start, i)
			} else if source[i] == '\'' { // missing check of prefixes
				tokenType = Constant
				i = getString(source, start, i)
			}
		} else if c == '/' && source[i+1] == '/' { // Find // comments
			i = bytes.Index(source[i+1:], []byte("\n"))
			if i == -1 {
				i = end
			}
			continue
		} else if c == '/' && source[i+1] == '*' { // Find /* comments */
			i = bytes.Index(source[i+1:], []byte("*/")) + 2
			continue
		} else if byteInSlice(c, []byte(":+-<>&|*=")) {
			tokenType = Syntax
			i++
			newCh := source[i]
			if newCh == c && c != '>' {
				i++
			} else if c == '-' && newCh == '>' {
				i++
			} else if newCh == '=' {
				i++
			}
		} else if byteInSlice(c, []byte("()[]{}~!?^%;/.,")) {
			tokenType = Syntax
			i++
			if c == '.' && byteInSlice(source[i], numChars) {
				tokenType = Constant
				i++
				for byteInSlice(source[i], intOrFloatDigits) {
					i++
				}
				for _, suffix := range []byte("lLfF") {
					if suffix == source[i] {
						i++
						break
					}
				}
			}
		} else if byteInSlice(source[i], numChars) { // integer
			tokenType = Constant
			if c == '0' && byteInSlice(source[i+1], []byte("xX")) {
				i += 2
				for byteInSlice(source[i], hexDigits) {
					i++
				}
			} else {
				for byteInSlice(source[i], intOrFloatDigits2) {
					i++
				}
			}
			// Handle integer and float suffixes
			// todo implement suffixes
		} else if c == '"' {
			tokenType = Constant
			i = getString(source, start, i)
		} else if c == '\'' {
			tokenType = Constant
			i = getChar(source, start, i)
		} else if c == '#' {
			tokenType = Preprocessor
			gotIf := source[i] == '#' &&
				source[i+1] == 'i' &&
				source[i+2] == 'f' &&
				source[i+3] == ' ' &&
				source[i+4] == ' '
			if gotIf {
				countIfs++
			} else if string(source[i:i+6]) == "#endif" {
				countIfs--
				if countIfs == 0 {
					ignoreErrors = false
				}
			}

			for true {
				/*i1 := bytes.Index(source[i:], []byte("\n"))
				i2 := bytes.Index(source[i:], []byte("//"))
				i3 := bytes.Index(source[i:], []byte("/*"))
				i4 := bytes.Index(source[i:], []byte("/"))*/

			}
		} else if ignoreErrors {
			fmt.Println("Dunno")
		} else {
			fmt.Println("Error")
		}

		if i <= 0 {
			fmt.Println("Invalid index, exit")
			return tokenSlice
		}

		//fmt.Println(token{tokenType, string(source[start:i]), start, i})
		tokenSlice = append(tokenSlice, &token{tokenType, string(source[start:i]), start, i})
	}
	return tokenSlice
}
