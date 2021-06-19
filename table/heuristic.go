// Copyright 2021 Rohit Awate
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package table

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
