package main

import (
	"fmt"

	"awate.in/commaql/commaql/parser"
)

func main() {
	query := "select distinct * from prices where id > 129 and name = 'soap'"
	// SELECT
	// STAR
	// FROM
	// identifier

	tokenizer := parser.Tokenizer{Query: query}
	tokens := tokenizer.Run()

	for _, token := range tokens {
		fmt.Println(token)
	}
}
