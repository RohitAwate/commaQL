package parser

import (
	"fmt"

	"awate.in/commaql/compiler/ast"
	"awate.in/commaql/compiler/parser/tokenizer"
)

func (p *Parser) selectColumnsList() []string {
	// TODO: Parse expressions

	var columnsList []string

	for p.match(tokenizer.IDENTIFIER) || p.match(tokenizer.STAR) {
		column := p.previous().Lexeme
		columnsList = append(columnsList, column)

		if !p.match(tokenizer.COMMA) {
			return columnsList
		}
	}

	p.emitError(fmt.Sprintf("Expected columns list, found %s", p.peek().Lexeme))
	return nil
}

func (p *Parser) selectTablesList() []string {
	var tablesList []string

	if p.match(tokenizer.IDENTIFIER) {
		table := p.previous().Lexeme
		tablesList = append(tablesList, table)
		return tablesList
	}

	p.emitError(fmt.Sprintf("Expected tables list, found %s", p.peek().Lexeme))
	return nil
}

func (p *Parser) whereClause() ast.Node {
	var expr ast.Expr
	if expr := p.expression(); expr == nil {
		return nil
	}

	for p.match(tokenizer.AND) || p.match(tokenizer.OR) {
		operator := p.previous()
		var rhs ast.Expr

		if rhs = p.expression().(ast.Expr); rhs == nil {
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
