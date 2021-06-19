package parser

import (
	"fmt"

	"awate.in/commaql/compiler/ast"
	"awate.in/commaql/compiler/parser/tokenizer"
)

func (p *Parser) selectStatement() ast.Node {
	// TODO: SELECT function()

	selectStmt := ast.SelectStmt{}

	var columns []string
	if columns = p.selectColumnsList(); columns == nil {
		p.emitError(fmt.Sprintf("Expected columns, found %s", p.peek().Lexeme))
		return nil
	}
	selectStmt.Columns = columns

	if !p.match(tokenizer.FROM) {
		p.emitError(fmt.Sprintf("Expected 'FROM', found %s", p.peek().Lexeme))
		return nil
	}

	var tables []string
	if tables := p.selectTablesList(); tables == nil {
		p.emitError(fmt.Sprintf("Expected tables list, found %s", p.peek().Lexeme))
		return nil
	}
	selectStmt.Tables = tables

	if p.match(tokenizer.WHERE) {
		whereClause := p.whereClause()
		if whereClause == nil {
			p.emitError(fmt.Sprintf("Expected expression, found %s", p.peek().Lexeme))
			return nil
		}

		selectStmt.WhereClause = whereClause.(ast.Expr)
	}

	if p.match(tokenizer.LIMIT) {
		limitExpr := p.expression()
		if limitExpr == nil {
			p.emitError(fmt.Sprintf("Expected limit expression, found %s", p.peek().Lexeme))
			return nil
		}

		selectStmt.Limit = limitExpr.(ast.Expr)
	}

	return selectStmt
}
