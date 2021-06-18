package parser

import (
	"fmt"

	"awate.in/commaql/compiler/parser/tokenizer"
)

func (p *Parser) selectColumnsList() bool {
	// TODO: Parse expressions

	for p.match(tokenizer.IDENTIFIER) || p.match(tokenizer.STAR) {
		if !p.match(tokenizer.COMMA) {
			return true
		}
	}

	p.emitError(fmt.Sprintf("Expected columns list, found %s", p.peek().Lexeme))
	return false
}

func (p *Parser) selectTablesList() bool {
	// TODO: Parse joins and shizz
	return p.match(tokenizer.IDENTIFIER)
}

func (p *Parser) whereClause() bool {
	if !p.expression() {
		return false
	}

	for p.match(tokenizer.AND) || p.match(tokenizer.OR) {
		if !p.expression() {
			return false
		}
	}

	return true
}
