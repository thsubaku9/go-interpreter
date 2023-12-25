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
	l.readChar()
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

func (l *Lexer) moveBack() {
	l.position, l.readPosition = l.readPosition-2, l.readPosition-1
	l.ch = l.input[l.readPosition]

	if l.readPosition == 0 {
		l.ch = 0
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
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

func (l *Lexer) NextToken() token.Token {
	defer l.readChar()
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			return token.Token{Type: token.EQ, Literal: literal}
		}
		return newToken(token.ASSIGN, l.ch)
	case ';':
		return newToken(token.SEMICOLON, l.ch)
	case '(':
		return newToken(token.LPAREN, l.ch)
	case ')':
		return newToken(token.RPAREN, l.ch)
	case ',':
		return newToken(token.COMMA, l.ch)
	case '+':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			return token.Token{Type: token.INC_BY, Literal: literal}
		}
		return newToken(token.PLUS, l.ch)
	case '-':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			return token.Token{Type: token.DEC_BY, Literal: literal}
		}
		return newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			return token.Token{Type: token.NOT_EQ, Literal: literal}
		}
		return newToken(token.BANG, l.ch)
	case '/':
		return newToken(token.SLASH, l.ch)
	case '*':
		return newToken(token.ASTERISK, l.ch)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			return token.Token{Type: token.LTE, Literal: literal}
		}
		return newToken(token.LT, l.ch)
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			return token.Token{Type: token.GTE, Literal: literal}
		}
		return newToken(token.GT, l.ch)
	case '{':
		return newToken(token.LBRACE, l.ch)
	case '}':
		return newToken(token.RBRACE, l.ch)
	case 0:
		return token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			var identifier string = l.readIdentifier()
			l.moveBack()
			return token.Token{Type: token.LookupIdent(identifier), Literal: identifier}
		} else if isDigit(l.ch) {
			var num string = l.readNumber()
			l.moveBack()
			return token.Token{Type: token.INT, Literal: num}
		}
	}

	return newToken(token.ILLEGAL, l.ch)
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
