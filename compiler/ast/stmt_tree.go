package ast

type SelectStmt struct {
	Columns       []string
	Tables        []string
	Limit         int
	GroupByClause GroupByClause
	OrderByClause OrderByClause
}

func (ss SelectStmt) amNode() {}
func (ss SelectStmt) amStmt() {}
