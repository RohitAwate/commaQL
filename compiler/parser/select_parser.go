package parser

import (
	"fmt"

	"awate.in/commaql/compiler/ast"
	"awate.in/commaql/compiler/parser/tokenizer"
)

func (p *Parser) selectStatement() ast.Node {
	// TODO: SELECT function()

	var columns []string

	if columns = p.selectColumnsList(); columns == nil {
		p.emitError(fmt.Sprintf("Expected columns, found %s", p.peek().Lexeme))
		return nil
	}

	if !p.match(tokenizer.FROM) {
		p.emitError(fmt.Sprintf("Expected 'FROM', found %s", p.peek().Lexeme))
		return nil
	}

	var tables []string

	if tables := p.selectTablesList(); tables == nil {
		p.emitError(fmt.Sprintf("Expected tables list, found %s", p.peek().Lexeme))
		return nil
	}

	if p.match(tokenizer.WHERE) {
		return p.whereClause()
	}

	return ast.SelectStmt{
		Columns: columns,
		Tables:  tables,
	}
}
