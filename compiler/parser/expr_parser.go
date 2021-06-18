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
	var expr ast.Expr

	if expr := p.logicalAND(); expr == nil {
		return nil
	}

	for p.match(tokenizer.OR) {
		operator := p.previous()
		var rhs ast.Expr

		if rhs := p.logicalAND(); rhs != nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr,
			Operator:     operator,
			RightOperand: rhs,
		}
	}

	return expr
}

func (p *Parser) logicalAND() ast.Node {
	var expr ast.Expr

	if expr := p.equality(); expr == nil {
		return nil
	}

	for p.match(tokenizer.AND) {
		operator := p.previous()
		var rhs ast.Expr

		if rhs := p.equality(); rhs != nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr,
			Operator:     operator,
			RightOperand: rhs,
		}
	}

	return expr
}

func (p *Parser) equality() ast.Node {
	var expr ast.Expr

	if expr := p.comparison(); expr == nil {
		return nil
	}

	for p.match(tokenizer.EQUALS) || p.match(tokenizer.NOT_EQUALS) {
		operator := p.previous()
		var rhs ast.Expr

		if rhs := p.comparison(); rhs != nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr,
			Operator:     operator,
			RightOperand: rhs,
		}
	}

	return expr
}

func (p *Parser) comparison() ast.Node {
	var expr ast.Expr

	if expr := p.term(); expr == nil {
		return nil
	}

	for p.match(tokenizer.GREATER_THAN) || p.match(tokenizer.LESS_THAN) ||
		p.match(tokenizer.GREATER_EQUALS) || p.match(tokenizer.LESS_EQUALS) {
		operator := p.previous()
		var rhs ast.Expr

		if rhs := p.term(); rhs != nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr,
			Operator:     operator,
			RightOperand: rhs,
		}
	}

	return expr
}

func (p *Parser) term() ast.Node {
	var expr ast.Expr

	if expr := p.factor(); expr == nil {
		return nil
	}

	for p.match(tokenizer.PLUS) || p.match(tokenizer.MINUS) {
		operator := p.previous()
		var rhs ast.Expr

		if rhs := p.factor(); rhs == nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr,
			Operator:     operator,
			RightOperand: rhs,
		}
	}

	return expr
}

func (p *Parser) factor() ast.Node {
	var expr ast.Expr

	if expr := p.exponent(); expr == nil {
		return nil
	}

	for p.match(tokenizer.STAR) || p.match(tokenizer.DIVIDE) || p.match(tokenizer.MODULO) {
		operator := p.previous()
		var rhs ast.Expr

		if rhs := p.exponent(); rhs != nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr,
			Operator:     operator,
			RightOperand: rhs,
		}
	}

	return expr
}

func (p *Parser) exponent() ast.Node {
	var expr ast.Expr

	if expr := p.unary(); expr == nil {
		return nil
	}

	for p.match(tokenizer.EXPONENT) {
		operator := p.previous()
		var rhs ast.Expr

		if rhs = p.comparison().(ast.Expr); rhs == nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr,
			Operator:     operator,
			RightOperand: rhs,
		}
	}

	return expr
}

func (p *Parser) unary() ast.Node {
	if p.match(tokenizer.MINUS) || p.match(tokenizer.NOT) {
		var expr ast.Expr

		if expr := p.unary(); expr == nil {
			return nil
		}

		return ast.UnaryExpr{
			Operator: p.previous(),
			Operand:  expr,
		}
	}

	return p.literal()
}

func (p *Parser) literal() ast.Node {
	switch p.peek().Type {
	case tokenizer.IDENTIFIER:
		fallthrough
	case tokenizer.NUMBER:
		fallthrough
	case tokenizer.TRUE:
		fallthrough
	case tokenizer.FALSE:
		fallthrough
	case tokenizer.NULL:
		fallthrough
	case tokenizer.STRING:
		p.consume()
	case tokenizer.OPEN_PAREN:
		return p.grouping()
	}

	p.emitError(fmt.Sprintf("Unexpected token: %s", p.peek().Lexeme))
	return nil
}

func (p *Parser) grouping() ast.Node {
	// Consume the '('
	p.consume()

	var innerExpr ast.Expr

	if innerExpr := p.expression(); innerExpr == nil {
		p.emitError("Expected expression")
		return nil
	}

	if p.match(tokenizer.CLOSE_PAREN) {
		p.emitError("Expected ')'")
		return nil
	}

	return ast.GroupedExpr{InnerExpr: innerExpr}
}
