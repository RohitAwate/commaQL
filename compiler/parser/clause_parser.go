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

func (p *Parser) selectColumnsList() []ast.SelectColumnNode {
	var columnsList []ast.SelectColumnNode

	for p.match(tokenizer.IDENTIFIER) || p.match(tokenizer.STAR) {
		column := ast.SelectColumnNode{
			ColumnToken: p.previous(),
		}

		columnsList = append(columnsList, column)

		if !p.match(tokenizer.COMMA) {
			return columnsList
		}
	}

	return nil
}

func (p *Parser) selectTablesList() []ast.TableNode {
	var tablesList []ast.TableNode

	if p.match(tokenizer.IDENTIFIER) {
		table := ast.TableNode{
			TableToken: p.previous(),
		}

		tablesList = append(tablesList, table)
		return tablesList
	}

	return nil
}

func (p *Parser) orderByClause() *ast.OrderByClause {
	var orderByColumns []ast.ColumnForOrderByClause
	if orderByColumns = p.orderByList(); orderByColumns == nil {
		p.emitExpectedError("table identifier(s)")
		return nil
	}

	return &ast.OrderByClause{
		Columns: orderByColumns,
	}
}

func (p *Parser) orderByList() []ast.ColumnForOrderByClause {
	var orderByColumns []ast.ColumnForOrderByClause

	for p.match(tokenizer.IDENTIFIER) {
		column := ast.ColumnForOrderByClause{
			ColumnToken: p.previous(),
		}

		orderByColumns = append(orderByColumns, column)

		if !p.match(tokenizer.COMMA) {
			return orderByColumns
		}
	}

	return nil
}

func (p *Parser) whereClause() ast.Node {
	expr := p.expression()
	if expr == nil {
		return nil
	}

	for p.match(tokenizer.AND) || p.match(tokenizer.OR) {
		operator := p.previous()

		rhs := p.expression()
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
