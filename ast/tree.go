package ast

import "awate.in/commaql/parser"

type Node interface {
	// Have to resort to shit like this thanks to lack of a strict inheritance system in Go.
	// Translates to "Hey, I am a Node!" <sigh>
	//
	// Inspired by Go's own AST implementation.
	// https://github.com/golang/go/blob/e3cb3817049ca5e9d96543500b72117f6ca659b8/src/go/ast/ast.go#L33-L36
	// https://github.com/golang/go/blob/e3cb3817049ca5e9d96543500b72117f6ca659b8/src/go/ast/ast.go#L534-L559
	amNode()
}

type Stmt interface {
	Node
	amStmt()
}

type Expr interface {
	Node
	amExpr()
}

type Clause interface {
	Node
	amClause()
}

type Literal struct {
	Meta parser.Token
}

type UnaryExpr struct {
	Operator parser.Token
	Operand  Expr
}

type BinaryExpr struct {
	LeftOperand  Expr
	Operator     parser.Token
	RightOperand Expr
}

type GroupedExpr struct {
	InnerExpr Expr
}

type SelectStmt struct {
	Columns       []string
	Limit         int
	GroupByClause GroupByClause
	OrderByClause OrderByClause
}

type ColumnForOrderByClause struct {
	Name      string
	Ascending bool
}

type OrderByClause struct {
	Columns []ColumnForOrderByClause
}

type GroupByClause struct {
	Columns []string
}

func (obc *OrderByClause) amClause() {}
func (gbc *GroupByClause) amClause() {}
