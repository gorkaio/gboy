package cpu

import (
	"fmt"
	"strings"
)

// Dasm disassembles an instruction
func Dasm(data uint) string {
	if data == 0x00 {
		return "NOP"
	}
	panic(fmt.Sprintf("Unknown instruction %#02x", data))
}

// Asm assemblies an instruction
func Asm(data string) uint {
	clearData := strings.ToUpper(data)
	if clearData == "NOP" {
		return 0x00
	}
	panic(fmt.Sprintf("Unknown instruction %s", data))
}
