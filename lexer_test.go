package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeLexer(input string) *lexer {
	l := &lexer{
		name:      "testLexer",
		input:     input,
		items:     make(chan item),
		line:      1,
		startLine: 1,
	}

	return l
}

func TestAccept(t *testing.T) {
	l := makeLexer("abcd☺")

	testcaseTable := []struct {
		val  string
		want bool
	}{
		{val: "b", want: false},
		{val: "c", want: false},
		{val: "12", want: false},
		{val: "a", want: true},
		{val: "a", want: false},
		{val: "b", want: true},
		{val: "de", want: false},
		{val: "dc", want: true},
		{val: "abcd", want: true},
		{val: "d", want: false},
		{val: "☹️", want: false},
		{val: "☺", want: true},
		{val: "a", want: false},
	}

	for _, tt := range testcaseTable {
		got := l.accept(tt.val)
		assert.Equal(t, tt.want, got, "with val %s", tt.val)
	}
}
