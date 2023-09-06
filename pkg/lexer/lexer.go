// Package lexer implements a lexer for the Monkey programming language.
// This lexer only supports ASCII characters, not the full Unicode character set.
// In order to do that Lexer.ch would need to be rune instead of byte, and the way of reading characters
// would need to change (since characters could take up multiple bytes in Unicode).
package lexer

import "gdejong/interpreter-in-go/pkg/token"

type Lexer struct {
	input        string // source code to lex
	position     int    // points to the character in input we are currently looking at
	readPosition int    // points to the next character in input, so we can peek ahead
	ch           byte   //
}

func New(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

// readChar gives us the next character and advances our position in the input string.
func (l *Lexer) readChar() {
	// Check if we've reached the end of the input string.
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" character
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	// Placeholder implementation.
	return token.Token{}
}
