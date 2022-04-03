package table

import (
	"encoding/csv"
	"os"
	"strings"
)

type CSVTable[T SQLType] struct {
	name    string
	columns []Column[T]
	reader  *csv.Reader
}

func NewCSVTable[T SQLType](file *os.File) (*CSVTable[T], error) {
	table := CSVTable[T]{
		reader: csv.NewReader(file), name: file.Name(),
		columns: []Column[T]{},
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

		switch deduceTypeForColumn(dataValue) {
		case SqlInt:
			newColumn := Column[int]{Name: columnName}
			table.columns = append(table.columns, newColumn)
		case SqlFloat:
			newColumn := Column[float64]{Name: columnName}
		case SqlBool:
			newColumn := Column[bool]{Name: columnName}
		default:
			newColumn := Column[string]{Name: columnName}
		}
	}

	return &table, nil
}

func (ct CSVTable) Name() string {
	return ct.name
}

func (ct CSVTable) Columns() []Column[T] {
	return ct.columns
}

func (ct CSVTable) LoadData() {
	// TODO
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
