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

package compiler

import (
	"fmt"
	"github.com/RohitAwate/commaql/compiler/codegen"
	"github.com/RohitAwate/commaql/compiler/parser"
	"github.com/RohitAwate/commaql/compiler/parser/tokenizer"
	"github.com/RohitAwate/commaql/disassembler"
	"github.com/RohitAwate/commaql/table"
	"github.com/RohitAwate/commaql/vm"
	"os"
)

type Compiler struct {
	tableContext map[string]table.Table
}

func NewCompiler(filepath string) (*Compiler, error) {
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("file not found: %s", filepath)
	}

	// TODO: Deduce filetype and build appropriate table
	var csvTable table.Table
	csvTable, err = table.NewCSVTable(reader)
	if err != nil {
		return nil, fmt.Errorf("could not read from file: %s", filepath)
	}

	tableContext := map[string]table.Table{
		table.GetTableNameFromFile(filepath): csvTable,
	}
	return &Compiler{tableContext: tableContext}, nil
}

func (c *Compiler) Compile(query string) vm.Bytecode {
	// TODO: Maybe return cg.Code from here so that code can be fed into VM

	t := tokenizer.Tokenizer{Query: query}
	tokens, _ := t.Run()

	p := parser.Parser{Tokens: tokens}
	statements, _ := p.Run()

	cg, _ := codegen.NewCodeGenerator(statements, c.tableContext)
	cg.Run()

	if len(cg.Errors) > 0 {
		for _, err := range cg.Errors {
			fmt.Println(err.Message)
		}
	}

	disassembler.Disassemble(&cg.Code)

	return cg.Code
}
