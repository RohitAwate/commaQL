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
