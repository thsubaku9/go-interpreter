package lexer

import (
	"monkey-i/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	defer l.readChar()
	return l
}

func (l *Lexer) IsFin() bool {
	return l.readPosition >= len(l.input)
}
func (l *Lexer) readChar() {

	if l.IsFin() {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position, l.readPosition = l.readPosition, l.readPosition+1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) NextToken() token.Token {

	l.skipWhitespace()

	switch l.ch {
	case '=':
		defer l.readChar()
		return newToken(token.ASSIGN, l.ch)
	case ';':
		defer l.readChar()
		return newToken(token.SEMICOLON, l.ch)
	case '(':
		defer l.readChar()
		return newToken(token.LPAREN, l.ch)
	case ')':
		defer l.readChar()
		return newToken(token.RPAREN, l.ch)
	case ',':
		defer l.readChar()
		return newToken(token.COMMA, l.ch)
	case '+':
		defer l.readChar()
		return newToken(token.PLUS, l.ch)
	case '{':
		defer l.readChar()
		return newToken(token.LBRACE, l.ch)
	case '}':
		defer l.readChar()
		return newToken(token.RBRACE, l.ch)
	case 0:
		defer l.readChar()
		return token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			var identifier string = l.readIdentifier()
			return token.Token{Type: token.LookupIdent(identifier), Literal: identifier}
		} else if isDigit(l.ch) {
			var num string = l.readNumber()
			return token.Token{Type: token.INT, Literal: num}
		}
	}

	return newToken(token.ILLEGAL, l.ch)
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
