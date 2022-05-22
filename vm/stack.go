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

package vm

import (
	"commaql/vm/values"
)

type stack struct {
	meta []values.Value
}

func (st *stack) push(v values.Value) {
	st.meta = append(st.meta, v)
}

func (st *stack) peek() values.Value {
	if len(st.meta) > 0 {
		return st.meta[len(st.meta)-1]
	}

	return nil
}

func (st *stack) pop() values.Value {
	if len(st.meta) > 0 {
		topOfStack := st.meta[len(st.meta)-1]
		st.meta = st.meta[:len(st.meta)-1]
		return topOfStack
	}

	return nil
}
