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

package codegen

import (
	"awate.in/commaql/compiler"
	"awate.in/commaql/compiler/ast"
)

type CodeGenerator struct {
	Bytecode compiler.Bytecode
}

func (cg *CodeGenerator) visitSelectStmt(ss *ast.SelectStmt) {

}

func (cg *CodeGenerator) visitOrderByClause(obc *ast.OrderByClause) {

}

func (cg *CodeGenerator) visitGroupByClause(gbc *ast.GroupByClause) {

}

func (cg *CodeGenerator) visitLiteral(l *ast.Literal) {

}

func (cg *CodeGenerator) visitUnaryExpr(ue *ast.UnaryExpr) {

}

func (cg *CodeGenerator) visitBinaryExpr(be *ast.BinaryExpr) {

}

func (cg *CodeGenerator) visitGroupedExpr(ge *ast.GroupedExpr) {

}
