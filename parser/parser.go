package parser

import (
	"log"
	"monkey-i/ast"
	"monkey-i/lexer"
	"monkey-i/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errs      []error
}

func New(l *lexer.Lexer) Parser {
	p := Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken, p.peekToken = p.peekToken, p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		// todo -> modify this to add errs
		return false
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = make([]ast.Statement, 0)

	for !p.curTokenIs(token.EOF) {
		var stmt ast.Statement = p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		p.nextToken()
	}

	return prog
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {

	if !p.curTokenIs(token.LET) {
		log.Panicf("Let statement starting token should be let, instead is %v at %+v\n", p.curToken, p.curToken.Cursor)
	}

	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		log.Panicf("Let statement identifier expected at %+v\n", p.curToken.Cursor)
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		log.Panicf("Let statement assignment expected at %+v\n", p.curToken.Cursor)
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
