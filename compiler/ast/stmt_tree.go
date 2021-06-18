package ast

type SelectStmt struct {
	Columns       []string
	Tables        []string
	WhereClause   Expr
	Limit         int
	GroupByClause GroupByClause
	OrderByClause OrderByClause
}

func (ss SelectStmt) amNode() {}
func (ss SelectStmt) amStmt() {}
