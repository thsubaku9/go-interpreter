package lexer

import (
	"monkey-i/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	lineNum      int
	barNum       int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, lineNum: 1, barNum: 1}
	l.readChar()
	return l
}

func (l *Lexer) isFin() bool {
	return l.readPosition >= len(l.input)
}

func (l *Lexer) readChar() {
	if l.isFin() {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
		l.barNum += 1
	}
	l.position, l.readPosition = l.readPosition, l.readPosition+1
}

func (l *Lexer) moveBack() {
	l.position, l.readPosition = l.readPosition-2, l.readPosition-1
	l.barNum -= 1

	if l.readPosition <= 0 {
		l.readPosition = 0
		l.position = -1
		l.barNum = 1
		l.ch = 0
	}

	if !l.isFin() {
		l.ch = l.input[l.readPosition]
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
		if l.ch == '\n' {
			l.lineNum += 1
			l.barNum = 1
		}

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

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
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
			return token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		}
		return newToken(token.ASSIGN, l.ch, l.lineNum, l.barNum)
	case ';':
		return newToken(token.SEMICOLON, l.ch, l.lineNum, l.barNum)
	case '(':
		return newToken(token.LPAREN, l.ch, l.lineNum, l.barNum)
	case ')':
		return newToken(token.RPAREN, l.ch, l.lineNum, l.barNum)
	case ',':
		return newToken(token.COMMA, l.ch, l.lineNum, l.barNum)
	case '+':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			return token.Token{Type: token.INC_BY, Literal: string(ch) + string(l.ch), Cursor: token.LineBar{Line: uint(l.lineNum), Bar: uint(l.barNum)}}
		}
		return newToken(token.PLUS, l.ch, l.lineNum, l.barNum)
	case '-':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			return token.Token{Type: token.DEC_BY, Literal: string(ch) + string(l.ch), Cursor: token.LineBar{Line: uint(l.lineNum), Bar: uint(l.barNum)}}
		}
		return newToken(token.MINUS, l.ch, l.lineNum, l.barNum)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			return token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch), Cursor: token.LineBar{Line: uint(l.lineNum), Bar: uint(l.barNum)}}
		}
		return newToken(token.BANG, l.ch, l.lineNum, l.barNum)
	case '/':
		return newToken(token.SLASH, l.ch, l.lineNum, l.barNum)
	case '*':
		return newToken(token.ASTERISK, l.ch, l.lineNum, l.barNum)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			return token.Token{Type: token.LTE, Literal: string(ch) + string(l.ch), Cursor: token.LineBar{Line: uint(l.lineNum), Bar: uint(l.barNum)}}
		}
		return newToken(token.LT, l.ch, l.lineNum, l.barNum)
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			return token.Token{Type: token.GTE, Literal: string(ch) + string(l.ch), Cursor: token.LineBar{Line: uint(l.lineNum), Bar: uint(l.barNum)}}
		}
		return newToken(token.GT, l.ch, l.lineNum, l.barNum)
	case '{':
		return newToken(token.LBRACE, l.ch, l.lineNum, l.barNum)
	case '}':
		return newToken(token.RBRACE, l.ch, l.lineNum, l.barNum)
	case 0:
		return newToken(token.EOF, "", l.lineNum, l.barNum)
	case '"':
		return newToken(token.STRING, l.readString(), l.lineNum, l.barNum)
	default:
		if isLetter(l.ch) {
			var identifier string = l.readIdentifier()
			l.moveBack()
			return token.Token{Type: token.LookupIdent(identifier), Literal: identifier, Cursor: token.LineBar{Line: uint(l.lineNum), Bar: uint(l.barNum)}}
		} else if isDigit(l.ch) {
			var num string = l.readNumber()
			l.moveBack()
			return token.Token{Type: token.INT, Literal: num, Cursor: token.LineBar{Line: uint(l.lineNum), Bar: uint(l.barNum)}}
		}
	}

	return newToken(token.ILLEGAL, l.ch, l.lineNum, l.barNum)
}

func newToken(tokenType token.TokenType, ch interface{}, lineNum int, barNum int) token.Token {

	switch literal := ch.(type) {
	case byte:
		return token.Token{Type: tokenType, Literal: string(literal), Cursor: token.LineBar{Line: uint(lineNum), Bar: uint(barNum)}}
	case string:
		return token.Token{Type: tokenType, Literal: literal, Cursor: token.LineBar{Line: uint(lineNum), Bar: uint(barNum)}}
	default:
		panic("newToken got unsupported ch input")
	}
}
