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
	"strings"
)

// TODO: Maybe move this to compiler/common.go
// TODO: Rename all to use Go's CamelCase style
// TODO: Rename all to be prefixed by Token____ since Go throws
// all of these into the global namespace
const (
	AND common.TokenType = iota // SQL Keywords
	ALL
	AS
	ASC
	BETWEEN
	BY
	CHECK
	COUNT
	DESC
	DISTINCT
	FALSE
	FROM
	FULL
	GROUP
	HAVING
	IN
	INNER
	IS
	JOIN
	LEFT
	LIKE
	LIMIT
	NOT
	NULL
	OR
	ORDER
	OUTER
	RIGHT
	SELECT
	TOP
	TRUE
	UNION
	WHERE
	STAR // Punctuation
	COMMA
	SINGLE_QUOTE
	DOUBLE_QUOTE
	DOT
	OPEN_PAREN
	CLOSE_PAREN
	SEMICOLON
	EQUALS
	NOT_EQUALS
	PLUS // Operators
	MINUS
	DIVIDE
	MODULO
	LESS_THAN
	GREATER_THAN
	LESS_EQUALS
	GREATER_EQUALS
	EXPONENT
	STRING
	IDENTIFIER
	NUMBER
)

var SQLKeywordsToTokenType = map[string]common.TokenType{
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
	"FALSE":    FALSE,
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
	"TRUE":     TRUE,
	"UNION":    UNION,
	"WHERE":    WHERE,
}

func IsSQLKeyword(identifier string) bool {
	_, ok := SQLKeywordsToTokenType[strings.ToUpper(identifier)]
	return ok
}
