// Copyright 2021 Rohit Awate
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
