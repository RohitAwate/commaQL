package parser

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
