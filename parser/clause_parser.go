package parser

import "fmt"

func (p *Parser) selectColumnsList() bool {
	// TODO: Parse expressions

	for p.match(IDENTIFIER) || p.match(STAR) {
		if !p.match(COMMA) {
			return true
		}
	}

	p.emitError(fmt.Sprintf("Expected columns list, found %s", p.peek().Lexeme))
	return false
}

func (p *Parser) selectTablesList() bool {
	// TODO: Parse joins and shizz
	return p.match(IDENTIFIER)
}

func (p *Parser) whereClause() bool {
	if !p.expression() {
		return false
	}

	for p.match(AND) || p.match(OR) {
		if !p.expression() {
			return false
		}
	}

	return true
}
