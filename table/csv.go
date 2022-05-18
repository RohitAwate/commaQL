package table

import (
	"encoding/csv"
	"fmt"
	"github.com/RohitAwate/commaql/vm/values"
	"os"
	"strings"
)

type CSVTable struct {
	name    string
	columns []Column
	reader  *csv.Reader
}

func NewCSVTable(file *os.File) (*CSVTable, error) {
	table := CSVTable{
		reader: csv.NewReader(file), name: file.Name(),
		columns: []Column{},
	}

	// Scan the first row of the CSVs.
	// It may or may not be the header row.
	firstRow, err := table.nextRow()
	if err != nil {
		return nil, err
	}

	// We run some simple heuristics to determine if the first row is a header.
	// A header row is the one which contains the names of all columns.
	containsHeader := isHeaderRow(firstRow)

	// Now we need to determine the data type of each column.
	// This is done based on the first row of data. If there's a header row,
	// this would be the 2nd overall row in the CSV. Otherwise, it is the first row.
	var dataRow []string
	if containsHeader {
		dataRow, err = table.nextRow()
		if err != nil {
			return nil, err
		}
	} else {
		dataRow = firstRow
	}

	// We iterate over the data points and try to deduce their types.
	for index, dataValue := range dataRow {
		var columnName string
		if containsHeader {
			columnName = firstRow[index]
		} else {
			columnName = getColumnAlias(uint(index))
		}

		newColumn := Column{Name: columnName, Type: deduceTypeForColumn(dataValue)}
		table.columns = append(table.columns, newColumn)
	}

	return &table, nil
}

func (ct CSVTable) Name() string {
	return ct.name
}

func (ct CSVTable) Columns() []Column {
	return ct.columns
}

func (ct CSVTable) LoadData() {
	// TODO
}

func (ct CSVTable) NextRow() ([]values.Value, error) {
	return make([]values.Value, 10), nil
}

func (ct CSVTable) nextRow() ([]string, error) {
	row, err := ct.reader.Read()
	if err != nil {
		return nil, err
	}

	for index := range row {
		row[index] = strings.TrimSpace(row[index])
	}

	return row, nil
}

func (ct CSVTable) IndexOfColumn(colName string) (uint, error) {
	for idx, col := range ct.columns {
		if col.Name == colName {
			return uint(idx), nil
		}
	}

	return 0, fmt.Errorf("column %s not found in table %s", colName, ct.Name())
}
