package parser

import "fmt"

func (p *Parser) selectStatement() bool {
	// TODO: SELECT function()

	if !p.selectColumnsList() {
		return false
	}

	if !p.match(FROM) {
		p.emitError(fmt.Sprintf("Expected 'FROM', found %s", p.peek().Lexeme))
		return false
	}

	if !p.selectTablesList() {
		p.emitError(fmt.Sprintf("Expected tables list, found %s", p.peek().Lexeme))
		return false
	}

	if p.match(WHERE) {
		return p.whereClause()
	}

	return true
}

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
