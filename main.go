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
	query := "select net, gross FROM prices limit 10*2 where id > 1"

	tokenizer := tokenizer.Tokenizer{Query: query}
	tokens, errors := tokenizer.Run()

	for _, token := range tokens {
		fmt.Println(token)
	}

	for _, err := range errors {
		fmt.Println(err)
	}

	parser := parser.Parser{Tokens: tokens}
	statements, errors := parser.Run()

	prettyPrint(statements)

	for _, err := range errors {
		fmt.Println(err)
	}
}
