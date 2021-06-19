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

package table

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Column struct {
	Name string
	Type SQLType
}

type Table struct {
	Name    string
	Columns []Column

	ContainsHeader bool
}

func GetTableFromCSV(csvFile *os.File) (*Table, error) {
	csvReader := csv.NewReader(csvFile)

	firstRow, err := NextRow(csvReader)
	if err != nil {
		return nil, err
	}

	table := &Table{ContainsHeader: isHeaderRow(firstRow)}

	var dataRow []string
	if table.ContainsHeader {
		dataRow, err = NextRow(csvReader)
		if err != nil {
			return nil, err
		}
	} else {
		dataRow = firstRow
	}

	for index, dataValue := range dataRow {
		var columnName string
		if table.ContainsHeader {
			columnName = firstRow[index]
		} else {
			columnName = GetColumnAlias(uint(index))
		}

		newColumn := Column{
			Name: columnName,
			Type: DeduceTypeForColumn(dataValue),
		}

		table.Columns = append(table.Columns, newColumn)
	}

	return table, nil
}

func NextRow(csvReader *csv.Reader) ([]string, error) {
	row, err := csvReader.Read()
	if err != nil {
		return nil, err
	}

	for index := range row {
		row[index] = strings.TrimSpace(row[index])
	}

	return row, nil
}

func GetColumnAlias(index uint) string {
	return fmt.Sprintf("Col_%d", index)
}
