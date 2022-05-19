package vm

import "github.com/RohitAwate/commaql/table"

type tableContext struct {
	table           table.Table
	selectedColumns []bool
	selectedRows    []uint
}

func newTableContext(tab table.Table) tableContext {
	return tableContext{
		table:           tab,
		selectedColumns: make([]bool, len(tab.Columns())),
	}
}

func (tc *tableContext) isColumnSelected(colIdx OpCode) bool {
	return tc.selectedColumns[colIdx]
}

func (tc *tableContext) markColumnSelected(colIdx OpCode) {
	tc.selectedColumns[colIdx] = true
}

func (tc *tableContext) markRowSelected(rowIdx uint) {
	tc.selectedRows = append(tc.selectedRows, rowIdx)
}
