package parser

import "strings"

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

func (t *Tokenizer) Run() []Token {
	t.Reset()

	var tokens []Token = make([]Token, 0)

	for !t.eof() {
		t.skipWhitespace()

		if isDigit(t.peek()) {
			tokens = append(tokens, t.number())
		} else if isAlpha(t.peek()) {
			tokens = append(tokens, t.identifier())
		}
	}

	return tokens
}

func (t *Tokenizer) eof() bool {
	return t.lookahead >= uint(len(t.Query))
}

func (t *Tokenizer) peek() byte {
	if t.lookahead < uint(len(t.Query)) {
		return t.Query[t.lookahead]
	} else {
		return byte(0)
	}
}

func (t *Tokenizer) advance() {
	if t.lookahead < uint(len(t.Query)) {
		t.lookahead++
	}
}

func (t *Tokenizer) advanceWindow() {
	t.lookahead++
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
		case '\r':
		case '\t':
			t.advance()
		case '\n':
			t.advance()
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
		token.Type = IDENTIFER
	}

	return token
}
