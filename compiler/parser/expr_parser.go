package parser

import (
	"fmt"

	"awate.in/commaql/compiler/ast"
	"awate.in/commaql/compiler/parser/tokenizer"
)

func (p *Parser) expression() ast.Node {
	return p.logicalOR()
}

func (p *Parser) logicalOR() ast.Node {
	if !p.logicalAND() {
		return false
	}

	for p.match(tokenizer.OR) {
		if !p.logicalAND() {
			return false
		}
	}

	return true
}

func (p *Parser) logicalAND() ast.Node {
	if !p.equality() {
		return false
	}

	for p.match(tokenizer.AND) {
		if !p.equality() {
			return false
		}
	}

	return true
}

func (p *Parser) equality() ast.Node {
	if !p.comparison() {
		return false
	}

	for p.match(tokenizer.EQUALS) || p.match(tokenizer.NOT_EQUALS) {
		if !p.comparison() {
			return false
		}
	}

	return true
}

func (p *Parser) comparison() ast.Node {
	if !p.term() {
		return false
	}

	for p.match(tokenizer.GREATER_THAN) || p.match(tokenizer.LESS_THAN) ||
		p.match(tokenizer.GREATER_EQUALS) || p.match(tokenizer.LESS_EQUALS) {
		if !p.term() {
			return false
		}
	}

	return true
}

func (p *Parser) term() ast.Node {
	expr := p.factor()

	for p.match(tokenizer.PLUS) || p.match(tokenizer.MINUS) {
		operator := p.previous()
		rhs := p.factor()

		expr = ast.BinaryExpr{
			LeftOperand:  expr.(ast.Expr),
			Operator:     operator,
			RightOperand: rhs.(ast.Expr),
		}
	}

	return expr
}

func (p *Parser) factor() ast.Node {
	if !p.exponent() {
		return false
	}

	for p.match(tokenizer.STAR) || p.match(tokenizer.DIVIDE) || p.match(tokenizer.MODULO) {
		if !p.exponent() {
			return false
		}
	}

	return true
}

func (p *Parser) exponent() ast.Node {
	if !p.unary() {
		return false
	}

	for p.match(tokenizer.EQUALS) || p.match(tokenizer.NOT_EQUALS) {
		if !p.comparison() {
			return false
		}
	}

	return true
}

func (p *Parser) unary() ast.Node {
	if p.match(tokenizer.MINUS) || p.match(tokenizer.NOT) {
		return p.unary()
	}

	return p.literal()
}

func (p *Parser) literal() ast.Node {
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

func (p *Parser) grouping() ast.Node {
	// Consume the '('
	p.consume()

	if !p.expression() {
		return false
	}

	if !p.match(tokenizer.CLOSE_PAREN) {
		p.emitError("Expected ')'")
		return false
	}

	return true
}
