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

	"github.com/RohitAwate/commaql/compiler"
	"github.com/RohitAwate/commaql/compiler/ast"
	"github.com/RohitAwate/commaql/compiler/parser/tokenizer"
)

type Parser struct {
	Tokens []compiler.Token

	current uint
	errors  []compiler.Error
}

func (p *Parser) Run() ([]ast.Stmt, []compiler.Error) {
	var statements []ast.Stmt

	for !p.eof() {
		statement := p.statement()
		if statement == nil {
			break
		}

		statements = append(statements, statement.(ast.Stmt))
	}

	return statements, p.errors
}

func (p *Parser) statement() ast.Node {
	var statement ast.Node

	if p.match(tokenizer.SELECT) {
		statement = p.selectStatement()
	}

	if p.eof() || p.match(tokenizer.SEMICOLON) {
		return statement
	} else {
		p.emitExpectedError("semicolon")
	}

	return nil
}

func (p *Parser) peek() compiler.Token {
	if p.current < uint(len(p.Tokens)) {
		return p.Tokens[p.current]
	}

	return p.Tokens[len(p.Tokens)-1]
}

func (p *Parser) previous() compiler.Token {
	return p.Tokens[p.current-1]
}

func (p *Parser) advance() {
	if p.current < uint(len(p.Tokens)) {
		p.current++
	}
}

func (p *Parser) match(tokenType compiler.TokenType) bool {
	if p.peek().Type == tokenType {
		p.advance()
		return true
	}

	return false
}

func (p *Parser) consume() {
	p.advance()
}

func (p *Parser) eof() bool {
	return p.current >= uint(len(p.Tokens))
}

func (p *Parser) emitError(msg string) {
	p.errors = append(p.errors, compiler.Error{Message: msg, Location: p.peek().Location})
}

func (p *Parser) emitExpectedError(expectedWhat string) {
	p.emitError(fmt.Sprintf("Expected %s, found '%s'", expectedWhat, p.peek().Lexeme))	
}