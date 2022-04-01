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
		if offset == 9 {
			fmt.Println()
		}

		opCode := bc.Blob[offset]
		opCodeInfo := vm.GetOpCodeInfo(opCode)
		var bytecodeLine strings.Builder
		bytecodeLine.WriteString(fmt.Sprintf("%5d\t%s%s%-15s%s", offset, ColorYellow, FormatBold, opCodeInfo.PrintableName, ColorReset))

		var argsInfoBuilder strings.Builder
		for argOffset := 1; argOffset < int(opCodeInfo.Args)+1; argOffset++ {
			arg := bc.Blob[offset+argOffset]
			bytecodeLine.WriteString(fmt.Sprintf(" %s%d%s", ColorBlue, arg, ColorReset))

			val := bc.ConstantsPool[arg]
			argsInfoBuilder.WriteString(fmt.Sprintf("\t %s# %d: %s%s\n", ColorBlue, arg, val, ColorReset))
		}

		fmt.Println(bytecodeLine.String())
		if opCodeInfo.Args > 0 {
			fmt.Print(argsInfoBuilder.String())
		}

		offset += int(opCodeInfo.Args)
	}
}
