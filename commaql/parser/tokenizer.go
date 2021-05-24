package parser

import "fmt"

type TokenType uint

const (
	// SQL Keywords
	AND      = iota
	ALL      = iota
	AS       = iota
	ASC      = iota
	BETWEEN  = iota
	CHECK    = iota
	COUNT    = iota
	DESC     = iota
	DISTINCT = iota
	FROM     = iota
	FULL     = iota
	GROUP_BY = iota
	HAVING   = iota
	IN       = iota
	INNER    = iota
	IS       = iota
	JOIN     = iota
	LEFT     = iota
	LIKE     = iota
	LIMIT    = iota
	NOT      = iota
	NULL     = iota
	OR       = iota
	ORDER_BY = iota
	OUTER    = iota
	RIGHT    = iota
	TOP      = iota
	UNION    = iota
	WHERE    = iota

	// Punctuation
	STAR         = iota
	COMMA        = iota
	SINGLE_QUOTE = iota
	DOUBLE_QUOTE = iota
	DOT          = iota
	OPEN_PAREN   = iota
	CLOSE_PAREN  = iota
	SEMICOLON    = iota

	// Operators
	PLUS         = iota
	MINUS        = iota
	DIVIDE       = iota
	LESS_THAN    = iota
	GREATER_THAN = iota
	EXPONENT     = iota

	IDENTIFER = iota
)

type Location struct {
	Line   uint
	Column uint
}

type Token struct {
	Type     TokenType
	Lexeme   string
	Location Location
}

type Tokenizer struct {
	Query string
}

func (t *Tokenizer) Run() []Token {
	var tokens []Token = make([]Token, 0)

	for index, char := range t.Query {
		fmt.Printf("%d %c\n", index, char)
	}

	return tokens
}
