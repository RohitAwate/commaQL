package parser

import "awate.in/commaql/table"

type ParserError struct {
	Message  string
	Location Location
}

type Parser struct {
	Table  table.Table
	Tokens []Token

	current uint
	errors  []ParserError
}

func (p *Parser) Run() (bool, []ParserError) {
	for !p.eof() {
		return p.statement(), p.errors
	}

	return true, p.errors
}

func (p *Parser) statement() bool {
	if p.match(SELECT) {
		return p.selectStatement()
	}

	return false
}

func (p *Parser) peek() Token {
	return p.Tokens[p.current]
}

func (p *Parser) advance() {
	if p.current < uint(len(p.Tokens))-1 {
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

func (p *Parser) emitError(msg string) {
	p.errors = append(p.errors, ParserError{Message: msg, Location: p.peek().Location})
}
