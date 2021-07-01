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

package compiler

import "awate.in/commaql/vm/types"

const (
	OP_ADD = iota + 6969
	OP_SUBTRACT
	OP_MULTIPLY
	OP_DIVIDE
	OP_EXPONENT
)

type OpCode uint

type Bytecode struct {
	Blob          []OpCode
	ConstantsPool []types.Value
}

func (b *Bytecode) AddConstant(v types.Value) {

}
