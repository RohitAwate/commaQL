// Copyright 2021-22 Rohit Awate
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
	"github.com/RohitAwate/commaql/compiler/codegen"
	"github.com/RohitAwate/commaql/compiler/parser"
	"github.com/RohitAwate/commaql/compiler/parser/tokenizer"
	"github.com/RohitAwate/commaql/disassembler"
)

func main() {
	query := `SELECT
				customer_id,
				first_name,
				last_name,
				amount,
				payment_date
			FROM
				customer
			WHERE amount = 100 - 2
			ORDER BY payment_date, amount;`

	t := tokenizer.Tokenizer{Query: query}
	tokens, _ := t.Run()

	p := parser.Parser{Tokens: tokens}
	statements, _ := p.Run()

	cg, _ := codegen.NewCodeGenerator(statements)
	cg.Run()

	disassembler.Disassemble(&cg.Code)
}
