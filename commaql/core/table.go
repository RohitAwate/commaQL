package core

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Column struct {
	Name string
	Type SQLType
	Data []interface{}
}

type Table struct {
	Name    string
	Columns []Column
}

func GetTableFromCSV(csvFile *os.File) (*Table, error) {
	r := csv.NewReader(csvFile)

	firstRow, err := r.Read()
	if err == io.EOF {
		return nil, fmt.Errorf("Unexpected EOF while reading %s", csvFile.Name())
	}

	if isHeaderRow(firstRow) {
		fmt.Println("Header detected")
	}

	return nil, nil
}
