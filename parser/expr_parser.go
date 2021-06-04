package parser

import "fmt"

func (p *Parser) expression() bool {
	return p.logicalOR()
}

func (p *Parser) logicalOR() bool {
	if !p.logicalAND() {
		return false
	}

	for p.match(OR) {
		if !p.logicalAND() {
			return false
		}
	}

	return true
}

func (p *Parser) logicalAND() bool {
	if !p.equality() {
		return false
	}

	for p.match(AND) {
		if !p.equality() {
			return false
		}
	}

	return true
}

func (p *Parser) equality() bool {
	if !p.comparison() {
		return false
	}

	for p.match(EQUALS) || p.match(NOT_EQUALS) {
		if !p.comparison() {
			return false
		}
	}

	return true
}

func (p *Parser) comparison() bool {
	if !p.term() {
		return false
	}

	for p.match(GREATER_THAN) || p.match(LESS_THAN) ||
		p.match(GREATER_EQUALS) || p.match(LESS_EQUALS) {
		if !p.term() {
			return false
		}
	}

	return true
}

func (p *Parser) term() bool {
	if !p.factor() {
		return false
	}

	for p.match(PLUS) || p.match(MINUS) {
		if !p.factor() {
			return false
		}
	}

	return true
}

func (p *Parser) factor() bool {
	if !p.exponent() {
		return false
	}

	for p.match(STAR) || p.match(DIVIDE) || p.match(MODULO) {
		if !p.exponent() {
			return false
		}
	}

	return true
}

func (p *Parser) exponent() bool {
	if !p.unary() {
		return false
	}

	for p.match(EQUALS) || p.match(NOT_EQUALS) {
		if !p.comparison() {
			return false
		}
	}

	return true
}

func (p *Parser) unary() bool {
	if p.match(MINUS) || p.match(NOT) {
		return p.unary()
	}

	return p.literal()
}

func (p *Parser) literal() bool {
	switch p.peek().Type {
	case IDENTIFIER:
		fallthrough
	case NUMBER:
		fallthrough
	case TRUE:
		fallthrough
	case FALSE:
		fallthrough
	case NULL:
		fallthrough
	case STRING:
		p.consume()
	case OPEN_PAREN:
		return p.grouping()
	default:
		p.emitError(fmt.Sprintf("Unexpected token: %s", p.peek().Lexeme))
		return false
	}

	return true
}

func (p *Parser) grouping() bool {
	// Consume the '('
	p.consume()

	if !p.expression() {
		return false
	}

	if !p.match(CLOSE_PAREN) {
		p.emitError("Expected ')'")
		return false
	}

	return true
}
