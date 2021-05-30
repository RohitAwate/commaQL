package core

import (
	"net/mail"
	"strconv"
)

func isHeaderRow(row []string) bool {
	return areNonEmptyStrings(row) &&
		!containsEmails(row) &&
		!containsNumbers(row) &&
		!containsBooleans(row)
}

func areNonEmptyStrings(row []string) bool {
	for _, col := range row {
		if len(col) < 1 {
			return false
		}
	}

	return true
}

func containsNumbers(row []string) bool {
	for _, col := range row {
		_, err := strconv.Atoi(col)
		if err == nil {
			return true
		}

		_, err = strconv.ParseFloat(col, 64)
		if err == nil {
			return true
		}
	}

	return false
}

func containsBooleans(row []string) bool {
	for _, col := range row {
		_, err := strconv.ParseBool(col)
		if err == nil {
			return true
		}
	}

	return false
}

func containsEmails(row []string) bool {
	for _, col := range row {
		_, err := mail.ParseAddress(col)
		if err == nil {
			return true
		}
	}

	return false
}
