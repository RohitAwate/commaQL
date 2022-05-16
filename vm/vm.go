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
	"fmt"
	"github.com/RohitAwate/commaql/table"
	"github.com/RohitAwate/commaql/vm/values"
)

type VM struct {
	// Table Context Registers
	// 0: Main register (all queries are executed with this context)
	// 1: Working register for joins, sub-queries, etc
	tcr   [2]*table.Table
	stack stack
	ip    uint
}

func NewVM() VM {
	return VM{
		tcr:   [2]*table.Table{},
		stack: stack{meta: []*values.Value{}},
		ip:    0,
	}
}

func binaryOp(left, right *values.Value, opCode OpCode) values.Value {
	if leftNum, ok := (*left).(*values.Number); ok {
		if rightNum, ok := (*right).(*values.Number); ok {
			switch opCode {
			case OpAdd:
				return values.NewNumberFromValue(leftNum.Meta + rightNum.Meta)
			case OpSubtract:
				return values.NewNumberFromValue(leftNum.Meta - rightNum.Meta)
			case OpMultiply:
				return values.NewNumberFromValue(leftNum.Meta * rightNum.Meta)
			case OpDivide:
				return values.NewNumberFromValue(leftNum.Meta / rightNum.Meta)
			case OpEquals:
				return values.NewBooleanFromValue(leftNum.Meta == rightNum.Meta)
			case OpNotEquals:
				return values.NewBooleanFromValue(leftNum.Meta != rightNum.Meta)
			case OpGreaterEquals:
				return values.NewBooleanFromValue(leftNum.Meta >= rightNum.Meta)
			case OpGreaterThan:
				return values.NewBooleanFromValue(leftNum.Meta > rightNum.Meta)
			case OpLessEquals:
				return values.NewBooleanFromValue(leftNum.Meta <= rightNum.Meta)
			case OpLessThan:
				return values.NewBooleanFromValue(leftNum.Meta < rightNum.Meta)
			default:
				panic("Not implemented!")
			}
		}
	}

	if leftVal, ok := (*left).(*values.Boolean); ok {
		if rightVal, ok := (*right).(*values.Boolean); ok {
			switch opCode {
			case OpAnd:
				return values.NewBooleanFromValue(leftVal.Meta && rightVal.Meta)
			case OpOr:
				return values.NewBooleanFromValue(leftVal.Meta || rightVal.Meta)
			default:
				panic("Not implemented!")
			}
		}
	}

	panic(0) // TODO: Handle this better! Also think about string concat
}

func (vm *VM) Run(bc Bytecode) {
	for vm.ip = 0; vm.ip < uint(len(bc.Blob)); vm.ip++ {
		opCode := bc.Blob[vm.ip]
		switch opCode {
		case OpLoadConst:
			vm.ip++
			constOffset := bc.Blob[vm.ip]
			vm.stack.push(&bc.ConstantsPool[constOffset])
		case OpAdd:
			fallthrough
		case OpSubtract:
			fallthrough
		case OpMultiply:
			fallthrough
		case OpDivide:
			fallthrough
		case OpAnd:
			fallthrough
		case OpOr:
			fallthrough
		case OpGreaterThan:
			fallthrough
		case OpGreaterEquals:
			fallthrough
		case OpLessThan:
			fallthrough
		case OpLessEquals:
			fallthrough
		case OpEquals:
			fallthrough
		case OpNotEquals:
			leftOperand := vm.stack.pop()
			rightOperand := vm.stack.pop()
			res := binaryOp(leftOperand, rightOperand, opCode)
			vm.stack.push(&res)
		case OpLoadTable:
			// OpLoadTable table_ctx_idx table_ctx_register_idx
			vm.ip++
			tableCtx := bc.TableContext[bc.Blob[vm.ip]]

			vm.ip++
			vm.tcr[bc.Blob[vm.ip]] = tableCtx
		default:
			panic("Instruction not implemented: " + string(opCode))
		}
	}

	// TODO: Get rid of this
	for _, val := range vm.stack.meta {
		fmt.Println(*val)
	}
}
