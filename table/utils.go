package table

import (
	"fmt"
	"strings"
)

func getColumnAlias(index uint) string {
	return fmt.Sprintf("Col_%d", index)
}

func GetTableNameFromFile(filename string) string {
	// TODO: Make this much, much smarter
	// - Handle multiple extensions: data.txt.csv > data
	// - Handle fully qualified paths: /home/rohit/project/data.csv > data

	dotLocation := strings.LastIndex(filename, ".")
	return filename[:dotLocation]
}
