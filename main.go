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
	"fmt"
	"github.com/RohitAwate/commaql/compiler"
	"github.com/RohitAwate/commaql/vm"
)

func main() {
	query := `select name, age from people where age >= 23 and graduated = false;`

	c, _ := compiler.NewCompiler("people.csv")
	bytecode := c.Compile(query)

	vm_ := vm.NewVM()
	resultSet := vm_.Run(bytecode)

	for _, row := range resultSet.Meta {
		fmt.Println(row)
	}
}
