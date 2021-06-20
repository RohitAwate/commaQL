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

package ast

import "awate.in/commaql/compiler"

type Literal struct {
	Meta compiler.Token
}

type UnaryExpr struct {
	Operator compiler.Token
	Operand  Expr
}

type BinaryExpr struct {
	LeftOperand  Expr
	Operator     compiler.Token
	RightOperand Expr
}

type GroupedExpr struct {
	InnerExpr Expr
}

func (l Literal) amNode()      {}
func (ue UnaryExpr) amNode()   {}
func (be BinaryExpr) amNode()  {}
func (ge GroupedExpr) amNode() {}

func (l Literal) amExpr()      {}
func (ue UnaryExpr) amExpr()   {}
func (be BinaryExpr) amExpr()  {}
func (ge GroupedExpr) amExpr() {}
