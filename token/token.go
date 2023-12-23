package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    LinePoint
}

type LinePoint struct {
	Row string
	Col string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals + constants
	IDENT = "IDENT"
	INT   = "INT"
	CONST = "CONST"
	// Operators
	ASSIGN = "="
	PLUS   = "+"
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
)
