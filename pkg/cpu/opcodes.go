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
	XOR_A = 0xAF
	LD_HL_D16 = 0x21
	LD_C_D8 = 0x0E
	LD_B_D8 = 0x06
	LDD_HL_A = 0x32
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
		fmt.Printf("NOP\n")
		return 4
	},
	JMP: func(cpu *CPU) int {
		addressLow := uint16(cpu.memory.Read(cpu.Registers.PC + 1))
		addressHigh := uint16(cpu.memory.Read(cpu.Registers.PC + 2))
		cpu.Registers.PC = addressHigh << 8 | addressLow
		fmt.Printf("JMP %#04x\n", cpu.Registers.PC)
		return 16
	},
	XOR_A: func(cpu *CPU) int {
		cpu.Registers.A = 0
		cpu.Registers.F |= 0x80
		cpu.Registers.PC++
		fmt.Printf("XOR A\n")
		return 4
	},
	LD_HL_D16: func(cpu *CPU) int {
		cpu.Registers.L = cpu.memory.Read(cpu.Registers.PC + 1)
		cpu.Registers.H = cpu.memory.Read(cpu.Registers.PC + 2)
		fmt.Printf("LD HL, %#02x\n", uint16(cpu.Registers.H) << 8 | uint16(cpu.Registers.L))
		cpu.Registers.PC+=3
		return 12
	},
	LD_C_D8: func(cpu *CPU) int {
		cpu.Registers.C = cpu.memory.Read(cpu.Registers.PC + 1)
		cpu.Registers.PC+=2
		fmt.Printf("LD C, %#02x\n", cpu.Registers.C)
		return 8
	},
	LD_B_D8: func(cpu *CPU) int {
		cpu.Registers.B = cpu.memory.Read(cpu.Registers.PC + 1)
		cpu.Registers.PC+=2
		fmt.Printf("LD B, %#02x\n", cpu.Registers.B)
		return 8
	},
	LDD_HL_A: func(cpu *CPU) int {
		cpu.memory.Write(uint16(cpu.Registers.H) << 8 | uint16(cpu.Registers.L), cpu.Registers.A)
		cpu.Registers.L--
		cpu.Registers.PC++
		fmt.Printf("LD (HL), A\n")
		return 8
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
