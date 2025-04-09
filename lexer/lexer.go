package lexer

import "josiahLang/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char, Lexer.ch)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Prepares Lexer to start read
	return l
}

func (l *Lexer) readChar() {
	// NOTE: Doesn't use a for loop for separation of concerns
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF or read nothing
	} else {
		l.ch = l.input[l.readPosition] // set ch to the next character
	}

	l.position = l.readPosition // point to position last read
	l.readPosition += 1         // always point to position next
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar() // ensures next NextToken() the l.ch field is already updated

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
