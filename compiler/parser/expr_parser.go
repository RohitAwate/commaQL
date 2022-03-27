// Copyright 2021-22 Rohit Awate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

import (
	"fmt"

	"github.com/RohitAwate/commaql/compiler/ast"
	"github.com/RohitAwate/commaql/compiler/parser/tokenizer"
)

func (p *Parser) expression() ast.Node {
	return p.logicalOR()
}

func (p *Parser) logicalOR() ast.Node {
	expr := p.logicalAND()
	if expr == nil {
		return nil
	}

	for p.match(tokenizer.OR) {
		operator := p.previous()

		rhs := p.logicalAND()
		if rhs == nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr.(ast.Expr),
			Operator:     operator,
			RightOperand: rhs.(ast.Expr),
		}
	}

	return expr
}

func (p *Parser) logicalAND() ast.Node {
	expr := p.equality()
	if expr == nil {
		return nil
	}

	for p.match(tokenizer.AND) {
		operator := p.previous()

		rhs := p.equality()
		if rhs == nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr.(ast.Expr),
			Operator:     operator,
			RightOperand: rhs.(ast.Expr),
		}
	}

	return expr
}

func (p *Parser) equality() ast.Node {
	expr := p.comparison()
	if expr == nil {
		return nil
	}

	for p.match(tokenizer.EQUALS) || p.match(tokenizer.NOT_EQUALS) {
		operator := p.previous()

		rhs := p.comparison()
		if rhs == nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr.(ast.Expr),
			Operator:     operator,
			RightOperand: rhs.(ast.Expr),
		}
	}

	return expr
}

func (p *Parser) comparison() ast.Node {
	expr := p.term()
	if expr == nil {
		return nil
	}

	for p.match(tokenizer.GREATER_THAN) || p.match(tokenizer.LESS_THAN) ||
		p.match(tokenizer.GREATER_EQUALS) || p.match(tokenizer.LESS_EQUALS) {
		operator := p.previous()

		rhs := p.term()
		if rhs == nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr.(ast.Expr),
			Operator:     operator,
			RightOperand: rhs.(ast.Expr),
		}
	}

	return expr
}

func (p *Parser) term() ast.Node {
	expr := p.factor()
	if expr == nil {
		return nil
	}

	for p.match(tokenizer.PLUS) || p.match(tokenizer.MINUS) {
		operator := p.previous()

		rhs := p.factor()
		if rhs == nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr.(ast.Expr),
			Operator:     operator,
			RightOperand: rhs.(ast.Expr),
		}
	}

	return expr
}

func (p *Parser) factor() ast.Node {
	expr := p.exponent()
	if expr == nil {
		return nil
	}

	for p.match(tokenizer.STAR) || p.match(tokenizer.DIVIDE) || p.match(tokenizer.MODULO) {
		operator := p.previous()

		rhs := p.exponent()
		if rhs == nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr.(ast.Expr),
			Operator:     operator,
			RightOperand: rhs.(ast.Expr),
		}
	}

	return expr
}

func (p *Parser) exponent() ast.Node {
	expr := p.unary()
	if expr == nil {
		return nil
	}

	for p.match(tokenizer.EXPONENT) {
		operator := p.previous()

		rhs := p.comparison()
		if rhs == nil {
			return nil
		}

		expr = ast.BinaryExpr{
			LeftOperand:  expr.(ast.Expr),
			Operator:     operator,
			RightOperand: rhs.(ast.Expr),
		}
	}

	return expr
}

func (p *Parser) unary() ast.Node {
	if p.match(tokenizer.MINUS) || p.match(tokenizer.NOT) {
		expr := p.unary()
		if expr == nil {
			return nil
		}

		return ast.UnaryExpr{
			Operator: p.previous(),
			Operand:  expr.(ast.Expr),
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

		return ast.Literal{
			Meta: p.previous(),
		}
	case tokenizer.OPEN_PAREN:
		return p.grouping()
	}

	p.emitError(fmt.Sprintf("Unexpected token: '%s'", p.peek().Lexeme))
	return nil
}

func (p *Parser) grouping() ast.Node {
	// Consume the '('
	p.consume()

	innerExpr := p.expression()
	if innerExpr == nil {
		// TODO: Print out location as well
		p.emitError("Expected expression")
		return nil
	}

	if p.match(tokenizer.CLOSE_PAREN) {
		// TODO: Print out location as well
		p.emitError("Expected ')'")
		return nil
	}

	return ast.GroupedExpr{InnerExpr: innerExpr.(ast.Expr)}
}
