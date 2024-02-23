package token

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Cursor  LineBar
}

type LineBar struct {
	Line uint
	Bar  uint
}

func (l LineBar) String() string {
	return fmt.Sprintf("%d:%d", l.Line, l.Bar)
}

const (
	// todo -> add support for and/or
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals + constants
	IDENT = "IDENT"
	INT   = "INT"
	CONST = "CONST"
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	// Two char ops
	EQ     = "=="
	NOT_EQ = "!="
	INC_BY = "+="
	DEC_BY = "-="
	LTE    = "<="
	GTE    = ">="
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	STRING   = "STRING"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"const":  CONST,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
