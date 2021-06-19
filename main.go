package main

import (
	"encoding/json"
	"fmt"

	"awate.in/commaql/compiler/parser"
	"awate.in/commaql/compiler/parser/tokenizer"
)

func prettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", " ")
	fmt.Println(string(s))
}

func main() {
	query := "select net, gross FROM prices where id > 10"

	tokenizer := tokenizer.Tokenizer{Query: query}
	tokens, errors := tokenizer.Run()

	for _, token := range tokens {
		fmt.Println(token)
	}

	for _, err := range errors {
		fmt.Println(err)
	}

	parser := parser.Parser{Tokens: tokens}
	tree, errors := parser.Run()

	prettyPrint(tree)

	for _, err := range errors {
		fmt.Println(err)
	}
}
