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
	"github.com/RohitAwate/commaql/compiler"
	"github.com/RohitAwate/commaql/compiler/ast"
	"github.com/RohitAwate/commaql/compiler/parser/tokenizer"
	"github.com/RohitAwate/commaql/vm"
	"github.com/RohitAwate/commaql/vm/types"
)

type CodeGenerator struct {
	statements []ast.Stmt
	Code       vm.Bytecode
	Errors     []compiler.Error
}

func NewCodeGenerator(statements []ast.Stmt) (*CodeGenerator, error) {
	if statements == nil {
		return nil, fmt.Errorf("root of AST cannot be nil")
	}

	return &CodeGenerator{
		statements: statements,
		Code:       vm.NewBytecode(),
	}, nil
}

func (cg *CodeGenerator) Run() compiler.PhaseStatus {
	for _, statement := range cg.statements {
		switch statement.(type) {
		case ast.SelectStmt:
			selectStmt := statement.(ast.SelectStmt)
			cg.visitSelectStmt(&selectStmt)
		}
	}

	return compiler.PHASE_OK
}

func (cg *CodeGenerator) visitSelectStmt(ss *ast.SelectStmt) {
	for _, tableNode := range ss.Tables {
		// Add table name to constants pool
		val := types.String{Meta: tableNode.TableToken.Lexeme}
		loc := cg.Code.AddConstant(val)

		// Load table name with LOAD_CONST
		cg.Code.EmitWithArg(vm.OpLoadConst, loc)
	}

	// SET_CTX
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
	switch e := interface{}(*expr).(type) {
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
		// TODO: Write a helper for this
		val := types.NewNumber(lit.Meta.Lexeme)
		loc := cg.Code.AddConstant(val)
		cg.Code.EmitWithArg(vm.OpLoadConst, loc)
	case tokenizer.TRUE:
		fallthrough
	case tokenizer.FALSE:
		val := types.NewBoolean(lit.Meta.Type)
		loc := cg.Code.AddConstant(val)
		cg.Code.EmitWithArg(vm.OpLoadConst, loc)
	case tokenizer.STRING:
		val := types.NewString(lit.Meta.Lexeme)
		loc := cg.Code.AddConstant(val)
		cg.Code.EmitWithArg(vm.OpLoadConst, loc)
	default:
		// FIXME
	}
}

func (cg *CodeGenerator) visitUnaryExpr(ue *ast.UnaryExpr) {

}

func (cg *CodeGenerator) visitBinaryExpr(be *ast.BinaryExpr) {
	cg.visitExpr(&be.RightOperand)
	cg.visitExpr(&be.LeftOperand)

	switch be.Operator.Type {
	case tokenizer.PLUS:
		cg.Code.Emit(vm.OpAdd)
	case tokenizer.MINUS:
		cg.Code.Emit(vm.OpSubtract)
	case tokenizer.STAR:
		cg.Code.Emit(vm.OpMultiply)
	case tokenizer.DIVIDE:
		cg.Code.Emit(vm.OpDivide)
	case tokenizer.MODULO:
		cg.Code.Emit(vm.OpModulo)
	case tokenizer.EXPONENT:
		cg.Code.Emit(vm.OpExponent)
	}
}

func (cg *CodeGenerator) visitGroupedExpr(ge *ast.GroupedExpr) {

}
