package cpu

import (
	"fmt"
	"strings"
	"errors"
	tablewriter "github.com/olekukonko/tablewriter"
	"os"
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
	DEC_B = 0x05
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
		printDebugInfo("NOP", cpu)
		return 4
	},
	JMP: func(cpu *CPU) int {
		addressLow := uint16(cpu.memory.Read(cpu.Registers.PC + 1))
		addressHigh := uint16(cpu.memory.Read(cpu.Registers.PC + 2))
		cpu.Registers.PC = addressHigh << 8 | addressLow
		printDebugInfo(fmt.Sprintf("JMP %#04x", cpu.Registers.PC), cpu)
		return 16
	},
	XOR_A: func(cpu *CPU) int {
		cpu.Registers.A = 0
		cpu.Registers.F |= 0x80
		cpu.Registers.PC++
		printDebugInfo("XOR A", cpu)
		return 4
	},
	LD_HL_D16: func(cpu *CPU) int {
		cpu.Registers.L = cpu.memory.Read(cpu.Registers.PC + 1)
		cpu.Registers.H = cpu.memory.Read(cpu.Registers.PC + 2)
		printDebugInfo(fmt.Sprintf("LD HL, %#02x", uint16(cpu.Registers.H) << 8 | uint16(cpu.Registers.L)), cpu)
		cpu.Registers.PC+=3
		return 12
	},
	LD_C_D8: func(cpu *CPU) int {
		cpu.Registers.C = cpu.memory.Read(cpu.Registers.PC + 1)
		cpu.Registers.PC+=2
		printDebugInfo(fmt.Sprintf("LD C, %#02x", cpu.Registers.C), cpu)
		return 8
	},
	LD_B_D8: func(cpu *CPU) int {
		cpu.Registers.B = cpu.memory.Read(cpu.Registers.PC + 1)
		cpu.Registers.PC+=2
		printDebugInfo(fmt.Sprintf("LD B, %#02x", cpu.Registers.B), cpu)
		return 8
	},
	LDD_HL_A: func(cpu *CPU) int {
		cpu.memory.Write(uint16(cpu.Registers.H) << 8 | uint16(cpu.Registers.L), cpu.Registers.A)
		cpu.Registers.L--
		cpu.Registers.PC++
		printDebugInfo("LD (HL), A", cpu)
		return 8
	},
	DEC_B: func(cpu *CPU) int {
		cpu.Registers.B--
		cpu.Registers.F |= 0x40 // Set N flag
		if cpu.Registers.B == 0 {
			cpu.Registers.F |= 0x80 // Set Z flag if zero
		}
		cpu.Registers.PC++
		printDebugInfo("DEC B", cpu)
		return 4
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

func printDebugInfo(operation string, cpu *CPU) {
	fmt.Printf("\n\n%s\n", operation)
	registerTable := tablewriter.NewWriter(os.Stdout)
	registerTable.SetHeader([]string{"Reg", "Value"})
	registerTable.Append([]string{"AF", fmt.Sprintf("%#04x", uint16(cpu.Registers.A) << 8 | uint16(cpu.Registers.F))})
	registerTable.Append([]string{"BC", fmt.Sprintf("%#04x", uint16(cpu.Registers.B) << 8 | uint16(cpu.Registers.C))})
	registerTable.Append([]string{"DE", fmt.Sprintf("%#04x", uint16(cpu.Registers.D) << 8 | uint16(cpu.Registers.E))})
	registerTable.Append([]string{"HL", fmt.Sprintf("%#04x", uint16(cpu.Registers.H) << 8 | uint16(cpu.Registers.L))})
	registerTable.Append([]string{"SP", fmt.Sprintf("%#04x", cpu.Registers.SP)})
	registerTable.Append([]string{"PC", fmt.Sprintf("%#04x", cpu.Registers.PC)})
	registerTable.Render()
	flagTable := tablewriter.NewWriter(os.Stdout)
	flagTable.SetHeader([]string{"Z", "N", "H", "C"})
	flagTable.Append([]string{
		fmt.Sprintf("%t", cpu.Registers.F & 0x80 > 1),
		fmt.Sprintf("%t", cpu.Registers.F & 0x40 > 1),
		fmt.Sprintf("%t", cpu.Registers.F & 0x20 > 1),
		fmt.Sprintf("%t", cpu.Registers.F & 0x10 > 1),
	})
	flagTable.Render()
}
