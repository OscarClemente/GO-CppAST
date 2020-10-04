package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type item struct {
	typ  itemType
	pos  int
	val  string
	line int
}

type itemType int

const (
	itemError itemType = iota
	itemChar
	itemString
	itemNumber
	itemName
	itemPreprocessor
	itemSyntax
)

const eof = -1

type stateFn func(*lexer) stateFn

type ILexer interface {
	next() rune
	peek() rune
	backup()
	emit(t itemType)
	ignore()
	accept(valid string) bool
	acceptRun(valid string) bool
	errorf(format string, args ...interface{}) stateFn
	nextItem() item
	drain()
}

// lexer holds the state of the lexing.
type lexer struct {
	name      string
	input     string
	pos       int
	start     int
	width     int
	items     chan item
	line      int
	startLine int
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}

	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width

	if r == '\n' {
		l.line++
	}

	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
	// Correct newline count if backedup rune was newline.
	if l.width == 1 && l.input[l.pos] == '\n' {
		l.line--
	}
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.start, l.input[l.start:l.pos], l.startLine}
	l.start = l.pos
	l.startLine = l.line
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) bool {
	existsAny := false
	for strings.ContainsRune(valid, l.next()) {
		existsAny = true
	}
	l.backup()
	return existsAny
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, l.start, fmt.Sprintf(format, args...), l.startLine}
	return nil
}

// nextItem returns the next item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) nextItem() item {
	return <-l.items
}

// drain drains the output so the lexing goroutine will exit.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) drain() {
	for range l.items {
	}
}
