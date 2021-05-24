package main

import (
	"fmt"

	"awate.in/commaql/commaql/parser"
)

func main() {
	query := "12345"
	// SELECT
	// STAR
	// FROM
	// identifier

	tokenizer := parser.Tokenizer{Query: query}
	tokens := tokenizer.Run()

	fmt.Println(tokens)
}
