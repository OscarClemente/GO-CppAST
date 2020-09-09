package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unicode"
)

func getString(source string, start int, i int) int {
	i = strings.Index(source[i+1:], "\"")
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
		i = strings.Index(source[i+1:], "\"")
	}
	return i + 1
}

func getChar(source string, start int, i int) int {
	i = strings.Index(source[i+1:], "'")
	for source[i-1] == '\\' {
		// Need special case '\\'
		if (i-2) > start && source[i-2] == '\\' {
			break
		}
		i = strings.Index(source[i+1:], "'")
	}
	// Unterminated single quotes
	if i < 0 {
		i = start
	}
	return i + 1
}

func isByteInString(b byte, s string) bool {
	// Is there really no bytes.ContainsByte()?
	return bytes.IndexByte([]byte(s), b) != -1
}

func isStringInStringSlice(referenceString string, ss []string) bool {
	for _, comparationString := range ss {
		if referenceString == comparationString {
			return true
		}
	}
	return false
}

//i = min([x for x in (i1, i2, i3, i4, end) if x != -1])
func minPositiveValue(iSlice []int) int {
	min := math.MaxInt32
	for _, val := range iSlice {
		if val >= 0 && val < min {
			min = val
		}
	}
	return min
}

func GetTokens(source string) []*token {
	letters := "abcdefghijklmnopqrstuvwxyz"
	numChars := "0123456789"
	extraChar := "_$"
	validChars := letters + strings.ToUpper(letters) + numChars + extraChar
	hexDigits := "0123456789abcdefABCDEF"
	intOrFloatDigits := "0123456789eE-+"
	intOrFloatDigits2 := "0123456789eE-+."
	strPrefixes := []string{"R", "u8", "u8R", "u", "uR", "U", "UR", "L", "LR"}

	fmt.Println(hexDigits)
	ignoreErrors := false
	countIfs := 0

	i := 0
	end := len(source)
	var tokenSlice []*token

	for i < end {
		// skip spaces
		for i < end && unicode.IsSpace(rune(source[i])) {
			i++
		}
		if i >= end {
			return tokenSlice
		}

		tokenType := Unknown
		start := i
		c := source[i]

		if unicode.IsLetter(rune(c)) || c == '_' {
			tokenType = Name
			for isByteInString(source[i], validChars) {
				i++
			}
			if source[i] == '\'' && (i-start) == 1 && strings.ContainsAny("uUL", source[start:i]) {
				//(source[start] == 'u' || source[start] == 'U' || source[start] == 'L') {
				// u, U and L are valid prefixes
				tokenType = Constant
				i = getChar(source, start, i)
			} else if source[i] == '\'' && isStringInStringSlice(source[start:i], strPrefixes) { // missing check of prefixes
				tokenType = Constant
				i = getString(source, start, i)
			}
		} else if c == '/' && source[i+1] == '/' { // Find // comments
			i = strings.Index(source[i+1:], "\n")
			if i == -1 {
				i = end
			}
			continue
		} else if c == '/' && source[i+1] == '*' { // Find /* comments */
			i = strings.Index(source[i+1:], "*/") + 2
			continue
		} else if isByteInString(c, ":+-<>&|*=") {
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
		} else if isByteInString(c, "()[]{}~!?^%;/.,") {
			tokenType = Syntax
			i++
			if c == '.' && isByteInString(source[i], numChars) {
				tokenType = Constant
				i++
				for isByteInString(source[i], intOrFloatDigits) {
					i++
				}
				if isByteInString(source[i], "lLfF") {
					i++
					break
				}
			}
		} else if isByteInString(source[i], numChars) { // integer
			tokenType = Constant
			if c == '0' && isByteInString(source[i+1], "xX") {
				i += 2
				for isByteInString(source[i], hexDigits) {
					i++
				}
			} else {
				for isByteInString(source[i], intOrFloatDigits2) {
					i++
				}
			}
			for _, suffix := range []string{"ull", "ll", "ul", "l", "f", "u"} {
				size := len(suffix)
				if suffix == strings.ToLower(source[i:i+size]) {
					i += size
					break
				}
			}
		} else if c == '"' {
			tokenType = Constant
			i = getString(source, start, i)
		} else if c == '\'' {
			tokenType = Constant
			i = getChar(source, start, i)
		} else if c == '#' {
			tokenType = Preprocessor
			gotIf := source[i:i+3] == "#if" && unicode.IsSpace(rune(source[i+3]))
			if gotIf {
				countIfs++
			} else if string(source[i:i+6]) == "#endif" {
				countIfs--
				if countIfs == 0 {
					ignoreErrors = false
				}
			}

			for true {
				i1 := strings.Index(source[i:], "\n")
				i2 := strings.Index(source[i:], "//")
				i3 := strings.Index(source[i:], "/*")
				i4 := strings.Index(source[i:], "/")

				i = minPositiveValue([]int{i1, i2, i3, i4}) + i

				if source[i] == '"' {
					i = strings.Index(source[i+1:], "\"") + 1
					if i > 0 {
						continue
					}
				}

				if !(i == i1 && source[i-1] == '\\') {
					if gotIf {
						condition := strings.TrimLeft(source[start+4:i], " ")
						if strings.HasPrefix(condition, "0") ||
							strings.HasPrefix(condition, "(0)") {
							ignoreErrors = true
						}
					}
					break
				}

				/*
				   if not (i == i1 and source[i-1] == '\\'):
				       if got_if:
				           condition = source[start+4:i].lstrip()
				           if (condition.startswith('0') or
				               condition.startswith('(0)')):
				               ignore_errors = True
				       break
				   i += 1*/
				i++
			}
		} else if c == '\\' {
			i++
			continue
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
