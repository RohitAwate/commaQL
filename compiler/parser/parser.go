package parser

import (
	"awate.in/commaql/compiler"
	"awate.in/commaql/compiler/parser/tokenizer"
	"awate.in/commaql/table"
)

type Parser struct {
	Table  table.Table
	Tokens []compiler.Token

	current uint
	errors  []compiler.Error
}

func (p *Parser) Run() (bool, []compiler.Error) {
	for !p.eof() {
		return p.statement(), p.errors
	}

	return true, p.errors
}

func (p *Parser) statement() bool {
	if p.match(tokenizer.SELECT) {
		return p.selectStatement()
	}

	return false
}

func (p *Parser) peek() compiler.Token {
	return p.Tokens[p.current]
}

func (p *Parser) previous() compiler.Token {
	return p.Tokens[p.current-1]
}

func (p *Parser) advance() {
	if p.current < uint(len(p.Tokens))-1 {
		p.current++
	}
}

func (p *Parser) match(tokenType compiler.TokenType) bool {
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
	p.errors = append(p.errors, compiler.Error{Message: msg, Location: p.peek().Location})
}
