package main

import (
	"bytes"
	"math"
	"strings"
)

// Constant definitions to help match some patterns
const letters string = "abcdefghijklmnopqrstuvwxyz"
const numChars string = "0123456789"
const extraChar string = "_$"
const validChars string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_$"
const hexDigits string = "0123456789abcdefABCDEF"
const intOrFloatDigits string = "0123456789eE-+"
const intOrFloatDigits2 string = "0123456789eE-+."

// findInString serves the purpose of returning the position of a substring
// inside a string whie being allowed to only search up from some point
// but returning the total index
func findInString(source string, fromPos int, substr string) int {
	i := strings.Index(source[fromPos:], substr)
	if i != -1 {
		i += fromPos
	}
	return i
}

// getString returns the end position of the string that is currently being read
// in most cases it will just find a " character and cut it at that
func getString(source string, start int, i int) int {
	i = findInString(source, i+1, "\"")
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
		i = findInString(source, i+1, "\"")
	}
	return i + 1
}

// getChar returns the position of the end of the char definition.
// Understanding a char definition as a 'c' including tildes
func getChar(source string, start int, i int) int {
	i = findInString(source, i+1, "'")
	for source[i-1] == '\\' {
		// Need special case '\\'
		if (i-2) > start && source[i-2] == '\\' {
			break
		}
		i = findInString(source, i+1, "'")
	}
	// Unterminated single quotes
	if i < 0 {
		i = start
	}
	return i + 1
}

// isByteInString returns true if the provided byte is one of the bytes in the string
func isByteInString(b byte, s string) bool {
	// Is there really no bytes.ContainsByte()?
	return bytes.IndexByte([]byte(s), b) != -1
}

// isStringInStringSlice returns true whether the provided string is equal to any of the
// strings inside the string slice
func isStringInStringSlice(referenceString string, ss []string) bool {
	for _, comparationString := range ss {
		if referenceString == comparationString {
			return true
		}
	}
	return false
}

// minPositiveValue returns the minimum value in the provided int slice that is above 0
func minPositiveValue(iSlice []int) int {
	min := math.MaxInt32
	for _, val := range iSlice {
		if val >= 0 && val < min {
			min = val
		}
	}
	return min
}
