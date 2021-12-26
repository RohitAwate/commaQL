package table

import (
	"fmt"
)

func getColumnAlias(index uint) string {
	return fmt.Sprintf("Col_%d", index)
}
