package main

import (
	"testing"
	"unicode/utf8"

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

// Testing next() function with general values.
func TestNext(t *testing.T) {
	testCases := []struct {
		input string
		want  rune
	}{
		{input: "a", want: 'a'},
		{input: "abc", want: 'a'},
		{input: "☺", want: '☺'},
		{input: "", want: eof},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			l := makeLexer(tc.input)
			got := l.next()

			runeLen := utf8.RuneLen(tc.want)
			if runeLen < 0 {
				runeLen = 0 // ignore pos if EOF
			}

			assert.Equal(t, tc.want, got)
			assert.Equal(t, runeLen, l.width)
			assert.Equal(t, runeLen, l.pos)
			assert.Equal(t, 1, l.line)
		})
	}
}

// Testing next() function with new line value,
// this is separated to streamline the general case
func TestNextNewLine(t *testing.T) {
	input := "\n"
	want := rune('\n')

	l := makeLexer(input)
	got := l.next()
	runeLen := utf8.RuneLen(want)

	assert.Equal(t, want, got)
	assert.Equal(t, l.width, runeLen)
	assert.Equal(t, l.pos, runeLen)
	assert.Equal(t, 2, l.line)
}
