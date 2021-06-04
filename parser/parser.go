package parser

import (
	"awate.in/commaql/core"
)

type ParserError struct {
	Message  string
	Location Location
}

type Parser struct {
	Table  core.Table
	Tokens []Token

	current uint
	errors  []ParserError
}

func (p *Parser) Run() bool {
	for !p.eof() {
		return p.statement()
	}

	return true
}

func (p *Parser) statement() bool {
	if p.match(SELECT) {
		return p.selectStatement()
	}

	return false
}

func (p *Parser) peek() Token {
	if p.current < uint(len(p.Tokens)) {
		return p.Tokens[p.current]
	}

	return Token{}
}

func (p *Parser) advance() {
	if p.current < uint(len(p.Tokens)) {
		p.current++
	}
}

func (p *Parser) match(tokenType TokenType) bool {
	if p.peek().Type == tokenType {
		p.advance()
		return true
	}

	return false
}

func (p *Parser) consume() {
	p.advance()
}

func (p *Parser) eof() bool {
	return p.current >= uint(len(p.Tokens))
}
