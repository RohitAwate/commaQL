package main

import "awate.in/commaql/commaql/parser"

func main() {
	query := "SELECT * FROM prices"
	// SELECT
	// STAR
	// FROM
	// identifier

	tokenizer := parser.Tokenizer{Query: query}
	tokenizer.Run()
}
