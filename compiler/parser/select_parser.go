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
	"commaql/compiler/ast"
	"commaql/compiler/parser/tokenizer"
)

func (p *Parser) selectStatement() ast.Node {
	selectStmt := ast.SelectStmt{}

	var columns []ast.SelectColumnNode
	if columns = p.selectColumnsList(); columns == nil {
		p.emitExpectedError("column identifier(s)")
		return nil
	}
	selectStmt.Columns = columns

	if !p.match(tokenizer.FROM) {
		p.emitExpectedError("'FROM'")
		return nil
	}

	var tables []ast.TableNode
	if tables = p.selectTablesList(); tables == nil {
		p.emitExpectedError("table identifier(s)")
		return nil
	}
	selectStmt.Tables = tables

	if p.match(tokenizer.WHERE) {
		whereClause := p.whereClause()
		if whereClause == nil {
			//p.emitExpectedError("expression")
			return nil
		}

		selectStmt.WhereClause = whereClause.(ast.Expr)
	}

	if p.match(tokenizer.ORDER) {
		if !p.match(tokenizer.BY) {
			p.emitExpectedError("'BY'")
			return nil
		}

		orderByClause := p.orderByClause()
		if orderByClause == nil {
			p.emitExpectedError("column identifier(s)")
			return nil
		}

		selectStmt.OrderByClause = *orderByClause
	}

	if p.match(tokenizer.LIMIT) {
		limitExpr := p.expression()
		if limitExpr == nil {
			p.emitExpectedError("expression")
			return nil
		}

		selectStmt.Limit = limitExpr.(ast.Expr)
	}

	return selectStmt
}
