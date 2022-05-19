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
	OpAnd
	OpOr
	OpNegate
	OpNot
	OpLoadTable
	OpJoin
	OpLoadConst

	OpScan

	OpSelectColumn
)

type OpCodeInfo struct {
	TableRegisterArgs  uint8
	ConstantOffsetArgs uint8
	InlineArgs         uint8
	StackArgs          uint8
	PrintableName      string
}

var opCodeInfoMap = map[OpCode]OpCodeInfo{
	OpAdd:           {0, 0, 0, 2, "OpAdd"},
	OpSubtract:      {0, 0, 0, 2, "OpSubtract"},
	OpMultiply:      {0, 0, 0, 2, "OpMultiply"},
	OpDivide:        {0, 0, 0, 2, "OpDivide"},
	OpModulo:        {0, 0, 0, 2, "OpModulo"},
	OpExponent:      {0, 0, 0, 2, "OpExponent"},
	OpGreaterThan:   {0, 0, 0, 2, "OpGreaterThan"},
	OpGreaterEquals: {0, 0, 0, 2, "OpGreaterEquals"},
	OpLessThan:      {0, 0, 0, 2, "OpLessThan"},
	OpLessEquals:    {0, 0, 0, 2, "OpLessEquals"},
	OpEquals:        {0, 0, 0, 2, "OpEquals"},
	OpNotEquals:     {0, 0, 0, 2, "OpNotEquals"},
	OpAnd:           {0, 0, 0, 2, "OpAnd"},
	OpOr:            {0, 0, 0, 2, "OpOr"},
	OpNegate:        {0, 0, 0, 1, "OpNegate"},
	OpNot:           {0, 0, 0, 1, "OpNot"},
	OpLoadTable:     {1, 0, 1, 0, "OpLoadTable"},
	OpJoin:          {0, 0, 0, 0, "OpJoin"},
	OpLoadConst:     {0, 1, 0, 0, "OpLoadConst"},
	OpSelectColumn:  {0, 0, 2, 0, "OpSelectColumn"},
}

func GetOpCodeInfo(opCode OpCode) OpCodeInfo {
	return opCodeInfoMap[opCode]
}

type Bytecode struct {
	Blob          []OpCode
	ConstantsPool []values.Value
	TableContext  []table.Table
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

func (b *Bytecode) AddTableContext(t table.Table) uint {
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
