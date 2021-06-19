package ast

type SelectStmt struct {
	Columns       []string
	Tables        []string
	WhereClause   Expr
	Limit         Expr
	GroupByClause GroupByClause
	OrderByClause OrderByClause
}

func (ss SelectStmt) amNode() {}
func (ss SelectStmt) amStmt() {}
