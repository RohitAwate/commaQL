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

package vm

import (
	"commaql/table"
	"commaql/vm/values"
)

type tableContext struct {
	table        table.Table
	selectedCols []uint
	selectedRows []uint
}

func newTableContext(tab table.Table) tableContext {
	return tableContext{table: tab}
}

func (tc *tableContext) markColumnSelected(colIdx OpCode) {
	tc.selectedCols = append(tc.selectedCols, uint(colIdx))
}

func (tc *tableContext) markRowSelected(rowIdx uint) {
	tc.selectedRows = append(tc.selectedRows, rowIdx)
}

func (tc *tableContext) toResultSet() ResultSet {
	result := ResultSet{
		Meta: make([][]values.Value, len(tc.selectedRows)),
	}

	for selRowIdx, selRowNum := range tc.selectedRows {
		result.Meta[selRowIdx] = make([]values.Value, len(tc.selectedCols))
		tabRow, _ := tc.table.GetRow(selRowNum)

		for selColIdx, selColNum := range tc.selectedCols {
			result.Meta[selRowIdx][selColIdx] = tabRow[selColNum]
		}
	}

	return result
}

type ResultSet struct {
	Meta [][]values.Value
}
