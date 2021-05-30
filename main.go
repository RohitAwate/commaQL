package main

import (
	"fmt"

	"awate.in/commaql/commaql/parser"
)

func main() {
	query := "SELECT WHERE 12345 FROM AS DISTINCT"
	// SELECT
	// STAR
	// FROM
	// identifier

	tokenizer := parser.Tokenizer{Query: query}
	tokens := tokenizer.Run()

	fmt.Println(tokens)
}
