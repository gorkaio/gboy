package cpu

import (
	"fmt"
	"strings"
	"errors"
)

const (
	unknownOpCode = "?"
	NOP = 0x00
	JMP = 0xC3
)

var opCodeToMnemonic = []string{
	"NOP",
}

var mnemonicToOpCode = map[string]int{
	"NOP": 0x00,
}

var opcodeHandlers = map[uint8]func(cpu *CPU) int {
	NOP: func(cpu *CPU) int {
		cpu.Registers.PC++
		return 4
	},
	JMP: func(cpu *CPU) int {
		addressLow := uint16(cpu.memory.Read(cpu.Registers.PC + 1))
		addressHigh := uint16(cpu.memory.Read(cpu.Registers.PC + 2))
		cpu.Registers.PC = addressHigh << 8 | addressLow
		fmt.Printf("JMP %#04x\n", cpu.Registers.PC)
		return 16
	},
}

// Dasm disassembles an instruction
func Dasm(data int) (string, error) {
	if (isOpCodeUnknown(data)) {
		errorString := fmt.Sprintf("Unknown instruction %#02x", data)
		return unknownOpCode, errors.New(errorString)
	}
	return opCodeToMnemonic[data], nil
}

// Asm assemblies an instruction
func Asm(data string) (int, error) {
	mnemonic := strings.ToUpper(data)
	if opcode, h := mnemonicToOpCode[mnemonic]; h {
		return opcode, nil
	}
	errorString := fmt.Sprintf("Unknown instruction %s", mnemonic)
	return 0, errors.New(errorString)
}

func isOpCodeUnknown(opcode int) bool {
	return len(opCodeToMnemonic) < opcode || opCodeToMnemonic[opcode] == unknownOpCode
}
