package table

import (
	"commaql/vm/values"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type CSVTable struct {
	name    string
	columns []ColumnInfo
	data    [][]values.Value
}

func parseRowOfValues(stringRow []string) ([]values.Value, error) {
	row := make([]values.Value, len(stringRow))

	for idx, val := range stringRow {
		switch deduceTypeForColumn(val) {
		case SqlInt:
			// TODO: Make a new values.Integer type
			fallthrough
		case SqlFloat:
			row[idx] = values.NewNumber(val)
		case SqlString:
			row[idx] = values.NewString(val)
		case SqlBool:
			row[idx] = values.NewBooleanFromString(val)
		default:
			panic("Type conversion not implemented!")
		}
	}

	return row, nil
}

func NewCSVTable(file *os.File) (*CSVTable, error) {
	reader := csv.NewReader(file)

	nextRow := func() ([]string, error) {
		row, err := reader.Read()
		if err != nil {
			return nil, err
		}

		for index := range row {
			row[index] = strings.TrimSpace(row[index])
		}

		return row, nil
	}

	table := CSVTable{
		name:    file.Name(),
		columns: []ColumnInfo{},
	}

	// Scan the first row of the CSVs.
	// It may or may not be the header row.
	firstRow, err := nextRow()
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
		dataRow, err = nextRow()
		if err != nil {
			return nil, err
		}
	} else {
		dataRow = firstRow
	}

	parsedRow, err := parseRowOfValues(dataRow)
	table.data = append(table.data, parsedRow)

	// We iterate over the data points and try to deduce their types.
	for index, dataValue := range dataRow {
		var columnName string
		if containsHeader {
			columnName = firstRow[index]
		} else {
			columnName = getColumnAlias(uint(index))
		}

		newColumn := ColumnInfo{Name: columnName, Type: deduceTypeForColumn(dataValue)}
		table.columns = append(table.columns, newColumn)
	}

	for {
		stringRow, err := nextRow()
		if err == io.EOF {
			break
		}

		parsedRow, err = parseRowOfValues(stringRow)
		table.data = append(table.data, parsedRow)
	}

	return &table, nil
}

func (ct CSVTable) Name() string {
	return ct.name
}

func (ct CSVTable) Columns() []ColumnInfo {
	return ct.columns
}

func (ct CSVTable) RowCount() uint {
	return uint(len(ct.data))
}

func (ct CSVTable) GetRow(rowIdx uint) ([]values.Value, error) {
	if rowIdx >= uint(len(ct.data)) {
		return nil, fmt.Errorf("row index out of range %d/%d", rowIdx, len(ct.data))
	}

	return ct.data[rowIdx], nil
}

func (ct CSVTable) IndexOfColumn(colName string) (uint, error) {
	for idx, col := range ct.columns {
		if col.Name == colName {
			return uint(idx), nil
		}
	}

	return 0, fmt.Errorf("column %s not found in table %s", colName, ct.Name())
}
