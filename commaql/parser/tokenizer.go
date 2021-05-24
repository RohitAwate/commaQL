package parser

type Tokenizer struct {
	Query string

	// Absolute trackers
	current   uint
	lookahead uint

	// Relative trackers
	currentLine   uint
	currentColumn uint

	lookaheadLine   uint
	lookaheadColumn uint
}

func (t *Tokenizer) Run() []Token {
	var tokens []Token = make([]Token, 0)

	for !t.eof() {
		if isDigit(t.peek()) {
			tokens = append(tokens, t.number())
		}
	}

	return tokens
}

func (t *Tokenizer) eof() bool {
	return t.lookahead == uint(len(t.Query))
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

func (t *Tokenizer) getLexemeForWindow() string {
	return t.Query[t.current:t.lookahead]
}

func (t *Tokenizer) getLocationForWindow() Location {
	return Location{Line: t.currentLine, Column: t.currentColumn}
}

func (t *Tokenizer) getTokenForWindow() Token {
	return Token{Lexeme: t.getLexemeForWindow(), Location: t.getLocationForWindow()}
}

func (t *Tokenizer) number() Token {
	for isDigit(t.peek()) {
		t.advance()
	}

	token := t.getTokenForWindow()
	token.Type = NUMBER

	return token
}
