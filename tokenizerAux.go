package main

import (
	"bytes"
	"math"
	"strings"
)

const letters string = "abcdefghijklmnopqrstuvwxyz"
const numChars string = "0123456789"
const extraChar string = "_$"
const validChars string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_$"
const hexDigits string = "0123456789abcdefABCDEF"
const intOrFloatDigits string = "0123456789eE-+"
const intOrFloatDigits2 string = "0123456789eE-+."

func findInString(source string, fromPos int, substr string) int {
	i := strings.Index(source[fromPos:], substr)
	if i != -1 {
		i += fromPos
	}
	return i
}

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

func minPositiveValue(iSlice []int) int {
	min := math.MaxInt32
	for _, val := range iSlice {
		if val >= 0 && val < min {
			min = val
		}
	}
	return min
}
