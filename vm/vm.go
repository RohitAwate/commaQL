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
	"commaql/vm/values"
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

	// Row register
	row []values.Value

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
			}
		}
	}

	if leftVal, ok := left.(values.String); ok {
		if rightVal, ok := right.(values.String); ok {
			switch opCode {
			case OpAdd:
				return values.NewString(leftVal.Meta + rightVal.Meta)
			case OpEquals:
				return values.NewBooleanFromValue(leftVal.Meta == rightVal.Meta)
			case OpNotEquals:
				return values.NewBooleanFromValue(leftVal.Meta != rightVal.Meta)
			}
		}
	}

	panic("UnsupportedOperation")
}

func (vm *VM) Run(bc Bytecode) ResultSet {
	readArg := func() OpCode {
		vm.ip++
		return bc.Blob[vm.ip]
	}

	for vm.ip = 0; vm.ip < uint(len(bc.Blob)); vm.ip++ {
		opCode := bc.Blob[vm.ip]

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
		case OpLoadNextRow:
			tab := vm.tcr[0].table
			vm.row, _ = tab.GetRow(vm.itr)
			vm.itr++
		case OpLoadVal:
			vm.stack.push(vm.row[readArg()])
		case OpJumpIfScan:
			// -1 to offset for the loop counter increment before the next iteration
			jumpOffset := uint(readArg()) - 1
			//fmt.Printf("Jumping to %s\n", GetOpCodeInfo(bc.Blob[jumpOffset+1]).PrintableName)
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

	return vm.tcr[0].toResultSet()
}
