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

package tokenizer

import (
	"commaql/compiler/common"
	"fmt"
	"strings"
)

type Tokenizer struct {
	Query string

	// Absolute trackers
	anchor    uint
	lookahead uint

	errors []common.Error
}

func (t *Tokenizer) Run() ([]common.Token, []common.Error) {
	t.Reset()

	tokens := make([]common.Token, 0)

	for !t.eof() {
		t.skipWhitespace()

		if isDigit(t.peek()) {
			tokens = append(tokens, t.number())
			continue
		} else if isAlpha(t.peek()) {
			tokens = append(tokens, t.identifier())
			continue
		}

		switch t.peek() {
		case '*':
			tokens = append(tokens, t.emitSingleCharToken(STAR))
		case ',':
			tokens = append(tokens, t.emitSingleCharToken(COMMA))
		case '\'':
			tokens = append(tokens, t.stringLiteral())
		case '"':
			// Consume opening quote
			t.consume()
			tokens = append(tokens, t.identifier())
			if t.peek() != '"' {
				t.emitError(fmt.Sprintf("Expected closing quote \", found '%c'", t.peek()))
				continue
			}

			// Consume closing quote
			t.consume()
		case '.':
			tokens = append(tokens, t.emitSingleCharToken(DOT))
		case '(':
			tokens = append(tokens, t.emitSingleCharToken(OPEN_PAREN))
		case ')':
			tokens = append(tokens, t.emitSingleCharToken(CLOSE_PAREN))
		case ';':
			tokens = append(tokens, t.emitSingleCharToken(SEMICOLON))
		case '=':
			tokens = append(tokens, t.emitSingleCharToken(EQUALS))
		case '+':
			tokens = append(tokens, t.emitSingleCharToken(PLUS))
		case '-':
			tokens = append(tokens, t.emitSingleCharToken(MINUS))
		case '/':
			tokens = append(tokens, t.emitSingleCharToken(DIVIDE))
		case '<':
			if t.peekNext() == '=' {
				t.advanceBy(2)
				token := t.emitToken()
				token.Type = LESS_EQUALS
				tokens = append(tokens, token)
				continue
			}

			tokens = append(tokens, t.emitSingleCharToken(LESS_THAN))
		case '>':
			if t.peekNext() == '=' {
				t.advanceBy(2)
				token := t.emitToken()
				token.Type = GREATER_EQUALS
				tokens = append(tokens, token)
				continue
			}

			tokens = append(tokens, t.emitSingleCharToken(GREATER_THAN))
		case '^':
			tokens = append(tokens, t.emitSingleCharToken(EXPONENT))
		case 0:
			// EOF
			break
		default:
			t.emitError(fmt.Sprintf("Unexpected token: '%c'", t.peek()))
		}
	}

	return tokens, t.errors
}

func (t *Tokenizer) Reset() {
	t.anchor = 0
	t.lookahead = 0
}

func (t *Tokenizer) emitError(msg string) {
	t.errors = append(t.errors, common.Error{Message: msg, Location: t.getLocationForWindow()})
}

func (t *Tokenizer) eof() bool {
	return t.lookahead >= uint(len(t.Query))
}

func (t *Tokenizer) peek() byte {
	if t.lookahead < uint(len(t.Query)) {
		// TODO: Try keeping just this, since this guard is already present in advance()
		return t.Query[t.lookahead]
	}

	return byte(0)
}

func (t *Tokenizer) peekNext() byte {
	if t.lookahead+1 < uint(len(t.Query)) {
		// TODO: Try keeping just this, since this guard is already present in advance()
		return t.Query[t.lookahead+1]
	}

	return byte(0)
}

func (t *Tokenizer) advanceBy(delta uint) {
	if t.lookahead+delta <= uint(len(t.Query)) {
		t.lookahead += delta
	}
}

func (t *Tokenizer) advance() {
	t.advanceBy(1)
}

func (t *Tokenizer) advanceWindow() {
	t.anchor = t.lookahead
}

func (t *Tokenizer) consume() {
	t.advance()
	t.advanceWindow()
}

func (t *Tokenizer) getLexemeForWindow() string {
	return t.Query[t.anchor:t.lookahead]
}

func (t *Tokenizer) getLocationForWindow() common.Location {
	// TODO: Track line and columns
	return common.Location{Line: 0, Column: 0}
}

func (t *Tokenizer) emitToken() common.Token {
	defer t.advanceWindow()
	return common.Token{Lexeme: t.getLexemeForWindow(), Location: t.getLocationForWindow()}
}

func (t *Tokenizer) skipWhitespace() {
	for {
		switch t.peek() {
		case ' ', '\r', '\t':
			t.advance()
			t.advanceWindow()
		case '\n':
			// TODO: Add line number tracking
			t.advance()
			t.advanceWindow()
		default:
			return
		}
	}
}

func (t *Tokenizer) number() common.Token {
	// TODO: Handle floats
	for isDigit(t.peek()) {
		t.advance()
	}

	token := t.emitToken()
	token.Type = NUMBER

	return token
}

func (t *Tokenizer) identifier() common.Token {
	for t.peek() == '_' || isAlpha(t.peek()) {
		t.advance()
	}

	token := t.emitToken()

	if IsSQLKeyword(token.Lexeme) {
		token.Type = SQLKeywordsToTokenType[strings.ToUpper(token.Lexeme)]
	} else {
		token.Type = IDENTIFIER
	}

	return token
}

func (t *Tokenizer) stringLiteral() common.Token {
	startingQuote := t.peek()

	// Consume opening quote
	t.consume()

	for t.peek() != startingQuote {
		t.advance()
	}

	stringToken := t.emitToken()
	stringToken.Type = STRING

	// Consume closing quote
	t.consume()

	return stringToken
}

func (t *Tokenizer) emitSingleCharToken(tokenType common.TokenType) common.Token {
	// TODO: Try and get rid of this and use just the emitToken method with a new
	// parameter that accept the token type. That would involve playing around with advance
	// since that appears to be handle by respective logic just before calling emitToken.
	// Might need another parameter that advances and then emits token.
	t.advance()
	token := t.emitToken()
	token.Type = tokenType
	return token
}
