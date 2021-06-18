package main

import (
	"fmt"

	"awate.in/commaql/compiler/parser"
	"awate.in/commaql/compiler/parser/tokenizer"
)

func main() {
	query := "select net, gross FROM prices where name > 10 and height <= 7 or age = 87 * 7"

	tokenizer := tokenizer.Tokenizer{Query: query}
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
