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

import "strconv"

type SQLType int
type SQLTypeSet interface {
	int | float64 | string | bool
}

const (
	SQL_INT    = iota
	SQL_FLOAT  = iota
	SQL_STRING = iota
	SQL_BOOL   = iota
)

func deduceTypeForColumn(sample string) SQLType {
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
