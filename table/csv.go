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
	table := &Table{ContainsHeader: isHeaderRow(firstRow)}

	// Now we need to determine the data type of each column.
	// This is done based on the first row of data. If there's a header row,
	// this would be the 2nd overall row in the CSV. Otherwise, it is the first row.
	var dataRow []string
	if table.ContainsHeader {
		dataRow, err = nextRow(csvReader)
		if err != nil {
			return nil, err
		}
	} else {
		dataRow = firstRow
	}

	// We iterate over the data points and try to deduce their types.
	for index, dataValue := range dataRow {
		var columnName string
		if table.ContainsHeader {
			columnName = firstRow[index]
		} else {
			columnName = getColumnAlias(uint(index))
		}

		newColumn := Column{
			Name: columnName,
			Type: deduceTypeForColumn(dataValue),
		}

		table.Columns = append(table.Columns, newColumn)
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
