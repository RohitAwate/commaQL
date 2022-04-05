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

import (
	"github.com/RohitAwate/commaql/table"
	"github.com/RohitAwate/commaql/vm/values"
)

type OpCode uint

const (
	OpAdd OpCode = iota + 6969
	OpSubtract
	OpMultiply
	OpDivide
	OpModulo
	OpExponent
	OpGreaterThan
	OpGreaterEquals
	OpLessThan
	OpLessEquals
	OpEquals
	OpNotEquals
	OpNegate
	OpNot
	OpLoadTable
	OpJoin
	OpLoadConst
	OpSetExecCtx
)

type OpCodeInfo struct {
	PrintableName string
	Args          uint8
}

var opCodeInfoMap = map[OpCode]OpCodeInfo{
	OpAdd:           {"OpAdd", 0},
	OpSubtract:      {"OpSubtract", 0},
	OpMultiply:      {"OpMultiply", 0},
	OpDivide:        {"OpDivide", 0},
	OpModulo:        {"OpModulo", 0},
	OpExponent:      {"OpExponent", 0},
	OpGreaterThan:   {"OpGreaterThan", 0},
	OpGreaterEquals: {"OpGreaterEquals", 0},
	OpLessThan:      {"OpLessThan", 0},
	OpLessEquals:    {"OpLessEquals", 0},
	OpEquals:        {"OpEquals", 0},
	OpNotEquals:     {"OpNotEquals", 0},
	OpNegate:        {"OpNegate", 0},
	OpNot:           {"OpNot", 0},
	OpLoadTable:     {"OpLoadTable", 0},
	OpJoin:          {"OpJoin", 0},
	OpLoadConst:     {"OpLoadConst", 1},
	OpSetExecCtx:    {"OpSetExecCtx", 0},
}

func GetOpCodeInfo(opCode OpCode) OpCodeInfo {
	return opCodeInfoMap[opCode]
}

type Bytecode struct {
	Blob          []OpCode
	ConstantsPool []values.Value
	TableContext  []*table.Table
}

func NewBytecode() Bytecode {
	return Bytecode{
		Blob:          []OpCode{},
		ConstantsPool: []values.Value{},
	}
}

func (b *Bytecode) AddConstant(v values.Value) uint {
	b.ConstantsPool = append(b.ConstantsPool, v)
	return uint(len(b.ConstantsPool) - 1)
}

func (b *Bytecode) AddTableContext(t *table.Table) uint {
	b.TableContext = append(b.TableContext, t)
	return uint(len(b.TableContext) - 1)
}

func (b *Bytecode) Emit(oc OpCode) uint {
	b.Blob = append(b.Blob, oc)
	return uint(len(b.Blob))
}

func (b *Bytecode) EmitWithArgs(oc OpCode, args ...uint) uint {
	convertedArgs := make([]OpCode, len(args)+1)

	convertedArgs[0] = oc
	for i, arg := range args {
		convertedArgs[i+1] = OpCode(arg)
	}

	b.Blob = append(b.Blob, convertedArgs...)
	return uint(len(b.Blob) - len(convertedArgs))
}
