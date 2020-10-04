package main

import "testing"

func TestAccept(t *testing.T) {
	l := &lexer{
		name:      "testLexer",
		input:     "abcdef",
		items:     make(chan item),
		line:      1,
		startLine: 1,
	}

	isAccepted := l.accept("a")
	if !isAccepted {
		t.Errorf("Expected: true, got: false")
	}
}
