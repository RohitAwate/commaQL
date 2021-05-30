package engine

import "strconv"

type SQLType int

const (
	SQL_INT    = iota
	SQL_FLOAT  = iota
	SQL_STRING = iota
	SQL_BOOL   = iota
)

func DeduceTypeForColumn(sample string) SQLType {
	_, err := strconv.Atoi(sample)
	if err == nil {
		return SQL_INT
	}

	_, err = strconv.ParseFloat(sample, 64)
	if err == nil {
		return SQL_FLOAT
	}

	_, err = strconv.ParseBool(sample)
	if err == nil {
		return SQL_BOOL
	}

	return SQL_STRING
}
