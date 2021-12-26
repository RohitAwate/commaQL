// Copyright 2021 Rohit Awate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"awate.in/commaql/compiler/parser"
	"awate.in/commaql/compiler/parser/tokenizer"
	"awate.in/commaql/table"
)

func prettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", " ")
	fmt.Println(string(s))
}

func main() {
	query := "select net, gross FROM prices where name = 'rohit' and \"type = 'card' limit 10*2 "

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

	file, _ := os.Open("imdb.csv")
	imdb, _ := table.GetTableFromCSV(file)
	fmt.Println(imdb)
}
