package parser

func (p *Parser) selectStatement() bool {
	// TODO: SELECT function()

	if !p.selectColumnsList() {
		return false
	}

	// TODO: return error if FROM not found
	if p.match(FROM) {
		return p.selectTablesList()
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

	return false
}

func (p *Parser) selectTablesList() bool {
	// TODO: Parse joins and shizz
	return p.match(IDENTIFIER)
}
