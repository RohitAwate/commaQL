package main

import (
	"fmt"
	"os"

	"awate.in/commaql/core"
	"awate.in/commaql/parser"
)

func main() {
	query := "select distinct * from prices where id > 129 and name = 'soap'"

	tokenizer := parser.Tokenizer{Query: query}
	tokens, errors := tokenizer.Run()

	for _, token := range tokens {
		fmt.Println(token)
	}

	for _, err := range errors {
		fmt.Println(err)
	}

	csvFile, _ := os.Open("prices.csv")

	table, err := core.GetTableFromCSV(csvFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(table)
}
