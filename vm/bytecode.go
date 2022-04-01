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

package vm

import "github.com/RohitAwate/commaql/vm/types"

type OpCode uint

const (
	OpAdd OpCode = iota + 6969
	OpSubtract
	OpMultiply
	OpDivide
	OpModulo
	OpExponent
	OpLoadTable
	OpJoin
	OpLoadConst
	OpSetExecCtx
)

type Bytecode struct {
	Blob          []OpCode
	ConstantsPool []types.Value
}

func NewBytecode() Bytecode {
	return Bytecode{
		Blob:          []OpCode{},
		ConstantsPool: []types.Value{},
	}
}

func (b *Bytecode) AddConstant(v types.Value) uint {
	b.ConstantsPool = append(b.ConstantsPool, v)
	return uint(len(b.ConstantsPool) - 1)
}

func (b *Bytecode) Emit(oc OpCode) uint {
	b.Blob = append(b.Blob, oc)
	return uint(len(b.Blob))
}

func (b *Bytecode) EmitWithArg(oc OpCode, arg uint) uint {
	b.Blob = append(b.Blob, oc, OpCode(arg))
	return uint(len(b.Blob) - 2)
}
