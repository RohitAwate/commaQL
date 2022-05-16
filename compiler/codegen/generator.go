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

package codegen

import (
	"fmt"

	"github.com/RohitAwate/commaql/compiler/ast"
	"github.com/RohitAwate/commaql/compiler/common"
	"github.com/RohitAwate/commaql/compiler/parser/tokenizer"
	"github.com/RohitAwate/commaql/table"
	"github.com/RohitAwate/commaql/vm"
	"github.com/RohitAwate/commaql/vm/values"
)

type CodeGenerator struct {
	statements   []ast.Stmt
	Code         vm.Bytecode
	Errors       []common.Error
	tableContext map[string]*table.Table
}

func NewCodeGenerator(statements []ast.Stmt, tableContext map[string]*table.Table) (*CodeGenerator, error) {
	if statements == nil {
		return nil, fmt.Errorf("root of AST cannot be nil")
	}

	return &CodeGenerator{
		statements:   statements,
		Code:         vm.NewBytecode(),
		tableContext: tableContext,
	}, nil
}

func (cg *CodeGenerator) Run() common.PhaseStatus {
	for _, statement := range cg.statements {
		switch stmt := statement.(type) {
		case ast.SelectStmt:
			cg.visitSelectStmt(&stmt)
		}
	}

	return common.PHASE_OK
}

func (cg *CodeGenerator) visitSelectStmt(ss *ast.SelectStmt) {
	for _, tableNode := range ss.Tables {
		if resolvedTable, ok := cg.tableContext[tableNode.TableToken.Lexeme]; !ok {
			cg.emitError(fmt.Sprintf("unknown table: %s", tableNode.TableToken.Lexeme), tableNode.TableToken)
			return
		} else {
			loc := cg.Code.AddTableContext(resolvedTable)
			var tableRegisterIndex uint = 0 // TODO: make this dynamic for a pair of tables at a time
			cg.Code.EmitWithArgs(vm.OpLoadTable, loc, tableRegisterIndex)
		}
	}

	cg.Code.Emit(vm.OpSetExecCtx)

	if ss.WhereClause != nil {
		cg.visitWhereClause(&ss.WhereClause)
	}
}

func (cg *CodeGenerator) visitWhereClause(expr *ast.Expr) {
	cg.visitExpr(expr)
}

func (cg *CodeGenerator) visitOrderByClause(obc *ast.OrderByClause) {

}

func (cg *CodeGenerator) visitGroupByClause(gbc *ast.GroupByClause) {

}

func (cg *CodeGenerator) visitExpr(expr *ast.Expr) {
	switch e := (*expr).(type) {
	case ast.UnaryExpr:
		cg.visitUnaryExpr(&e)
	case ast.BinaryExpr:
		cg.visitBinaryExpr(&e)
	case ast.GroupedExpr:
		cg.visitGroupedExpr(&e)
	case ast.Literal:
		cg.visitLiteral(&e)
	}
}

func (cg *CodeGenerator) visitLiteral(lit *ast.Literal) {
	switch lit.Meta.Type {
	case tokenizer.NUMBER:
		// TODO: Write a helper for this, lots of duplication between these cases
		val := values.NewNumber(lit.Meta.Lexeme)
		loc := cg.Code.AddConstant(val)
		cg.Code.EmitWithArgs(vm.OpLoadConst, loc)
	case tokenizer.TRUE:
		fallthrough
	case tokenizer.FALSE:
		val := values.NewBoolean(lit.Meta.Type)
		loc := cg.Code.AddConstant(val)
		cg.Code.EmitWithArgs(vm.OpLoadConst, loc)
	case tokenizer.STRING:
		val := values.NewString(lit.Meta.Lexeme)
		loc := cg.Code.AddConstant(val)
		cg.Code.EmitWithArgs(vm.OpLoadConst, loc)
	default:
		// FIXME
	}
}

var unaryOperatorToOpCode = map[common.TokenType]vm.OpCode{
	tokenizer.MINUS: vm.OpNegate,
	tokenizer.NOT:   vm.OpNot,
}

func (cg *CodeGenerator) visitUnaryExpr(ue *ast.UnaryExpr) {
	cg.visitExpr(&ue.Operand)
	cg.Code.Emit(unaryOperatorToOpCode[ue.Operator.Type])
}

var binaryOperatorToOpCode = map[common.TokenType]vm.OpCode{
	tokenizer.PLUS:           vm.OpAdd,
	tokenizer.MINUS:          vm.OpSubtract,
	tokenizer.STAR:           vm.OpMultiply,
	tokenizer.DIVIDE:         vm.OpDivide,
	tokenizer.MODULO:         vm.OpModulo,
	tokenizer.EXPONENT:       vm.OpExponent,
	tokenizer.GREATER_THAN:   vm.OpGreaterThan,
	tokenizer.GREATER_EQUALS: vm.OpGreaterEquals,
	tokenizer.LESS_THAN:      vm.OpLessThan,
	tokenizer.LESS_EQUALS:    vm.OpLessEquals,
	tokenizer.EQUALS:         vm.OpEquals,
	tokenizer.NOT_EQUALS:     vm.OpNotEquals,
	tokenizer.AND:            vm.OpAnd,
	tokenizer.OR:             vm.OpOr,
}

func (cg *CodeGenerator) visitBinaryExpr(be *ast.BinaryExpr) {
	cg.visitExpr(&be.RightOperand)
	cg.visitExpr(&be.LeftOperand)

	// TODO: Get rid of this
	if _, ok := binaryOperatorToOpCode[be.Operator.Type]; !ok {
		panic("Binary operator to opcode mapping not found: " + string(be.Operator.Lexeme))
	}

	cg.Code.Emit(binaryOperatorToOpCode[be.Operator.Type])
}

func (cg *CodeGenerator) visitGroupedExpr(ge *ast.GroupedExpr) {
	cg.visitExpr(&ge.InnerExpr)
}

func (cg *CodeGenerator) emitError(msg string, token common.Token) {
	cg.Errors = append(cg.Errors, common.Error{Message: msg, Location: token.Location})
}
