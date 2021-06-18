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