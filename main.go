package main

import (
	"fmt"

	"awate.in/commaql/parser"
)

func main() {
	query := "select net, gross FROM prices"

	tokenizer := parser.Tokenizer{Query: query}
	tokens, errors := tokenizer.Run()

	for _, token := range tokens {
		fmt.Println(token)
	}

	for _, err := range errors {
		fmt.Println(err)
	}

	parser := parser.Parser{Tokens: tokens}
	ok, errors := parser.Run()
	fmt.Printf("Parser run: %t\n", ok)

	for _, err := range errors {
		fmt.Println(err)
	}
}
