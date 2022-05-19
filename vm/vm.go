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
	"github.com/RohitAwate/commaql/vm/values"
)

type VM struct {
	// Table Context Registers
	// 0: Main register (all queries are executed with this context)
	// 1: Working register for joins, sub-queries, etc
	tcr [2]tableContext

	// Limit and Iterator Registers
	// Used for maintaining scan position within contexts.
	lim uint
	itr uint

	stack stack
	ip    uint
}

func NewVM() VM {
	return VM{
		tcr:   [2]tableContext{},
		stack: stack{meta: []values.Value{}},
		ip:    0,
	}
}

func binaryOp(left, right values.Value, opCode OpCode) values.Value {
	if leftNum, ok := left.(values.Number); ok {
		if rightNum, ok := right.(values.Number); ok {
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

	if leftVal, ok := left.(values.Boolean); ok {
		if rightVal, ok := right.(values.Boolean); ok {
			switch opCode {
			case OpAnd:
				return values.NewBooleanFromValue(leftVal.Meta && rightVal.Meta)
			case OpOr:
				return values.NewBooleanFromValue(leftVal.Meta || rightVal.Meta)
			case OpEquals:
				return values.NewBooleanFromValue(leftVal.Meta == rightVal.Meta)
			case OpNotEquals:
				return values.NewBooleanFromValue(leftVal.Meta != rightVal.Meta)
			default:
				panic("Not implemented!")
			}
		}
	}

	panic(0) // TODO: Handle this better! Also think about string concat
}

func (vm *VM) Run(bc Bytecode) {
	readArg := func() OpCode {
		vm.ip++
		return bc.Blob[vm.ip]
	}

	for vm.ip = 0; vm.ip < uint(len(bc.Blob)); vm.ip++ {
		opCode := bc.Blob[vm.ip]

		fmt.Printf("Executing: %s\n", GetOpCodeInfo(opCode).PrintableName)

		switch opCode {
		case OpLoadConst:
			vm.stack.push(bc.ConstantsPool[readArg()])
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
			vm.stack.push(res)
		case OpLoadTable:
			// OpLoadTable table_ctx_idx table_ctx_register_idx
			tableCtx := bc.TableContext[readArg()]
			vm.tcr[readArg()] = newTableContext(tableCtx)
		case OpSelectColumn:
			vm.tcr[readArg()].markColumnSelected(readArg())
		case OpScan:
			rows := vm.tcr[0].table.RowCount()
			vm.lim = rows
			vm.itr = 0
		case OpLoadVal:
			tab := vm.tcr[0].table
			row, _ := tab.NextRow()
			vm.stack.push(row[readArg()])
			vm.itr++
		case OpJumpIfScan:
			jumpOffset := uint(readArg())
			if vm.itr < vm.lim {
				vm.ip = jumpOffset
			} else {
				vm.itr = 0
				vm.lim = 0
			}
		case OpSelectRowIfTrue:
			if vm.stack.pop().(values.Boolean).Meta {
				vm.tcr[0].markRowSelected(vm.itr - 1)
			}
		default:
			panic("instruction not implemented: " + GetOpCodeInfo(opCode).PrintableName)
		}
	}

	// TODO: Get rid of this
	for _, val := range vm.stack.meta {
		fmt.Println(val)
	}

	for _, t := range vm.tcr {
		fmt.Println(t)
	}
}
