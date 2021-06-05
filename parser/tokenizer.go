package parser

import (
	"fmt"
	"strings"
)

type Tokenizer struct {
	Query string

	// Absolute trackers
	anchor    uint
	lookahead uint
}

func (t *Tokenizer) Reset() {
	t.anchor = 0
	t.lookahead = 0
}

func (t *Tokenizer) Run() ([]Token, []ParserError) {
	t.Reset()

	var tokens []Token = make([]Token, 0)
	var errors []ParserError = make([]ParserError, 0)

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
			tokens = append(tokens, t.emitSingleCharToken(SINGLE_QUOTE))
		case '"':
			tokens = append(tokens, t.emitSingleCharToken(DOUBLE_QUOTE))
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
		default:
			errMsg := fmt.Sprintf("Unexpected token: %c", t.peek())
			errors = append(errors, t.emitError(errMsg))
		}
	}

	return tokens, errors
}

func (t *Tokenizer) emitError(message string) ParserError {
	return ParserError{
		Message:  message,
		Location: t.getLocationForWindow(),
	}
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

func (t *Tokenizer) getLexemeForWindow() string {
	return t.Query[t.anchor:t.lookahead]
}

func (t *Tokenizer) getLocationForWindow() Location {
	// TODO: Track line and columns
	return Location{Line: 0, Column: 0}
}

func (t *Tokenizer) emitToken() Token {
	defer t.advanceWindow()
	return Token{Lexeme: t.getLexemeForWindow(), Location: t.getLocationForWindow()}
}

func (t *Tokenizer) skipWhitespace() {
	for {
		switch t.peek() {
		case ' ':
			fallthrough
		case '\r':
			fallthrough
		case '\t':
			t.advance()
			t.advanceWindow()
		case '\n':
			t.advance()
			t.advanceWindow()
		default:
			return
		}
	}
}

func (t *Tokenizer) number() Token {
	// TODO: Handle floats
	for isDigit(t.peek()) {
		t.advance()
	}

	token := t.emitToken()
	token.Type = NUMBER

	return token
}

func (t *Tokenizer) identifier() Token {
	for isAlpha(t.peek()) {
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

func (t *Tokenizer) emitSingleCharToken(tokenType TokenType) Token {
	t.advance()
	token := t.emitToken()
	token.Type = tokenType
	return token
}
