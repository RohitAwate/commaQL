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

	p.emitError(fmt.Sprintf("Expected columns list, found '%s'", p.peek().Lexeme))
	return nil
}

func (p *Parser) selectTablesList() []string {
	var tablesList []string

	if p.match(tokenizer.IDENTIFIER) {
		table := p.previous().Lexeme
		tablesList = append(tablesList, table)
		return tablesList
	}

	p.emitError(fmt.Sprintf("Expected tables list, found '%s'", p.peek().Lexeme))
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
