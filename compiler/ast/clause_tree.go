package ast

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

func (obc OrderByClause) amNode() {}
func (gbc GroupByClause) amNode() {}

func (obc OrderByClause) amClause() {}
func (gbc GroupByClause) amClause() {}
