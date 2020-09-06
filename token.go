package main

import (
	"fmt"
)

type TokenType string

const (
	Unknown      TokenType = "Unknown"
	Syntax                 = "Syntax"
	Constant               = "Constant"
	Name                   = "Name"
	Preprocessor           = "Preprocessor"
)

type token struct {
	tokenType TokenType
	name      string
	start     int
	end       int
}

func (t token) print() {
	fmt.Println(t.tokenType, t.name, t.start, t.end)
}
