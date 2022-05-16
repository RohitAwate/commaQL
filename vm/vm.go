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
	"fmt"
	"github.com/RohitAwate/commaql/vm/values"
)

type VM struct {
	ctx ExecutionContext

	stack stack
	ip    uint
}

func NewVM() VM {
	return VM{
		ctx:   ExecutionContext{},
		stack: stack{meta: []*values.Value{}},
		ip:    0,
	}
}

func (vm *VM) Run(bc Bytecode) {
	for vm.ip = 0; vm.ip < uint(len(bc.Blob)); vm.ip++ {
		switch bc.Blob[vm.ip] {
		case OpLoadConst:
			vm.ip++
			constOffset := bc.Blob[vm.ip]
			vm.stack.push(&bc.ConstantsPool[constOffset])
		}
	}

	fmt.Println(vm.stack.meta)
}
