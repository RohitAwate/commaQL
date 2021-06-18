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
