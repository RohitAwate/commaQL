// Copyright 2021-22 Rohit Awate
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

package types

import (
	"github.com/RohitAwate/commaql/compiler"
	"github.com/RohitAwate/commaql/compiler/parser/tokenizer"
	"strconv"
)

type Value interface {
	amValue()
}

type Number struct {
	Meta float64
}

func (n Number) amValue() {}

func NewNumber(lexeme string) *Number {
	number, _ := strconv.ParseFloat(lexeme, 64)
	return &Number{Meta: number}
}

type String struct {
	Meta string
}

func (n String) amValue() {}

func NewString(lexeme string) *String {
	return &String{Meta: lexeme}
}

type Boolean struct {
	Meta bool
}

func (b Boolean) amValue() {}

func NewBoolean(tokenType compiler.TokenType) *Boolean {
	return &Boolean{Meta: tokenType == tokenizer.TRUE}
}
