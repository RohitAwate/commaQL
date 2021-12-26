package table

import (
	"encoding/csv"
	"os"
	"strings"
)

func GetTableFromCSV(csvFile *os.File) (*Table, error) {
	csvReader := csv.NewReader(csvFile)

	// Scan the first row of the CSVs.
	// It may or may not be the header row.
	firstRow, err := nextRow(csvReader)
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
		dataRow, err = nextRow(csvReader)
		if err != nil {
			return nil, err
		}
	} else {
		dataRow = firstRow
	}

	var columns []Column[SQLTypeSet]
	// We iterate over the data points and try to deduce their types.
	for index, dataValue := range dataRow {
		var columnName string
		if containsHeader {
			columnName = firstRow[index]
		} else {
			columnName = getColumnAlias(uint(index))
		}

		switch deduceTypeForColumn(dataValue) {
		case SQL_INT:
			columns = append(columns, Column[int]{Name: columnName})
		case SQL_FLOAT:
			columns = append(columns, Column[float64]{Name: columnName})
		case SQL_BOOL:
			columns = append(columns, Column[bool]{Name: columnName})
		case SQL_STRING:
			columns = append(columns, Column[string]{Name: columnName})
		}
	}

	return table, nil
}

func nextRow(csvReader *csv.Reader) ([]string, error) {
	row, err := csvReader.Read()
	if err != nil {
		return nil, err
	}

	for index := range row {
		row[index] = strings.TrimSpace(row[index])
	}

	return row, nil
}
