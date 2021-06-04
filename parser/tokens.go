package parser

import (
	"strings"
)

type TokenType uint

const (
	// SQL Keywords
	AND      = iota
	ALL      = iota
	AS       = iota
	ASC      = iota
	BETWEEN  = iota
	BY       = iota
	CHECK    = iota
	COUNT    = iota
	DESC     = iota
	DISTINCT = iota
	FALSE    = iota
	FROM     = iota
	FULL     = iota
	GROUP    = iota
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
	ORDER    = iota
	OUTER    = iota
	RIGHT    = iota
	SELECT   = iota
	TOP      = iota
	TRUE     = iota
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
	EQUALS       = iota
	NOT_EQUALS   = iota

	// Operators
	PLUS           = iota
	MINUS          = iota
	DIVIDE         = iota
	MODULO         = iota
	LESS_THAN      = iota
	GREATER_THAN   = iota
	LESS_EQUALS    = iota
	GREATER_EQUALS = iota
	EXPONENT       = iota

	STRING     = iota
	IDENTIFIER = iota
	NUMBER     = iota
)

var SQLKeywordsToTokenType = map[string]TokenType{
	"AND":      AND,
	"ALL":      ALL,
	"AS":       AS,
	"ASC":      ASC,
	"BETWEEN":  BETWEEN,
	"BY":       BY,
	"CHECK":    CHECK,
	"COUNT":    COUNT,
	"DESC":     DESC,
	"DISTINCT": DISTINCT,
	"FROM":     FROM,
	"FULL":     FULL,
	"GROUP":    GROUP,
	"HAVING":   HAVING,
	"IN":       IN,
	"INNER":    INNER,
	"IS":       IS,
	"JOIN":     JOIN,
	"LEFT":     LEFT,
	"LIKE":     LIKE,
	"LIMIT":    LIMIT,
	"NOT":      NOT,
	"NULL":     NULL,
	"OR":       OR,
	"ORDER":    ORDER,
	"OUTER":    OUTER,
	"RIGHT":    RIGHT,
	"SELECT":   SELECT,
	"TOP":      TOP,
	"UNION":    UNION,
	"WHERE":    WHERE,
}

func IsSQLKeyword(identifier string) bool {
	for keyword := range SQLKeywordsToTokenType {
		if strings.ToUpper(identifier) == keyword {
			return true
		}
	}

	return false
}

type Location struct {
	Line   uint
	Column uint
}

type Token struct {
	Type     TokenType
	Lexeme   string
	Location Location
}
