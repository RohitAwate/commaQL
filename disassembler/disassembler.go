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

package disassembler

import (
	"fmt"
	"github.com/RohitAwate/commaql/vm"
	"strings"
)

const (
	ColorRed        = "\u001b[31m"
	ColorGreen      = "\u001b[32m"
	ColorYellow     = "\u001b[33m"
	ColorBlue       = "\u001b[34m"
	ColorWhite      = "\u001b[37m"
	ColorReset      = "\u001b[0m"
	FormatBold      = "\u001b[1m"
	FormatUnderline = "\u001b[4m"
)

func Disassemble(bc *vm.Bytecode) {
	for offset := 0; offset < len(bc.Blob); offset++ {
		opCode := bc.Blob[offset]
		opCodeInfo := vm.GetOpCodeInfo(opCode)

		var bytecodeLine strings.Builder
		bytecodeLine.WriteString(fmt.Sprintf("%5d\t%s%s%-15s%s", offset, ColorRed, FormatBold, opCodeInfo.PrintableName, ColorReset))

		for argOffset := 1; argOffset <= int(opCodeInfo.InlineArgs); argOffset++ {
			arg := bc.Blob[offset+argOffset]
			bytecodeLine.WriteString(fmt.Sprintf(" %si:%d%s", ColorGreen, arg, ColorReset))
		}

		var constOffsetArgs strings.Builder
		for argOffset := 1; argOffset <= int(opCodeInfo.ConstantOffsetArgs); argOffset++ {
			arg := bc.Blob[offset+argOffset]
			bytecodeLine.WriteString(fmt.Sprintf(" %sc:%d%s", ColorBlue, arg, ColorReset))

			val := bc.ConstantsPool[arg]
			constOffsetArgs.WriteString(fmt.Sprintf("\t %s# c:%d: %s%s\n", ColorBlue, arg, val, ColorReset))
		}

		for argOffset := 1; argOffset <= int(opCodeInfo.TableRegisterArgs); argOffset++ {
			arg := bc.Blob[offset+argOffset]
			bytecodeLine.WriteString(fmt.Sprintf(" %st:%d%s", ColorYellow, arg, ColorReset))
		}

		fmt.Println(bytecodeLine.String())
		if opCodeInfo.ConstantOffsetArgs > 0 {
			fmt.Print(constOffsetArgs.String())
		}

		offset += int(opCodeInfo.InlineArgs) + int(opCodeInfo.TableRegisterArgs) + int(opCodeInfo.ConstantOffsetArgs)
	}
}
