package main

import (
	"fmt"

	"awate.in/commaql/parser"
)

func main() {
	query := "select net, gross from prices"

	tokenizer := parser.Tokenizer{Query: query}
	tokens, errors := tokenizer.Run()

	for _, token := range tokens {
		fmt.Println(token)
	}

	for _, err := range errors {
		fmt.Println(err)
	}

	parser := parser.Parser{Tokens: tokens}
	ok := parser.Run()
	fmt.Printf("Parser run: %t\n", ok)
}
