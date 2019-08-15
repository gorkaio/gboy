package cpu

import (
	"errors"
	"fmt"

	memory "github.com/gorkaio/gboy/pkg/memory"
)

// Registers are CPU registeres
type Registers struct {
	A uint8
	F uint8
	B uint8
	C uint8
	H uint8
	L uint8
	PC uint16
}

// CPU structure
type CPU struct {
	Registers Registers
	memory    memory.MemoryInterface
}

// New initialises a new Z80 cpu
func New(memory memory.MemoryInterface) *CPU {
	return &CPU{
		Registers: Registers{
			PC: uint16(0x100),
		},
		memory: memory,
	}
}

// Step executes next instruction and returns cycles consumed
func (cpu *CPU) Step() (int, error) {
	opCode := cpu.memory.Read(cpu.Registers.PC)
	if handler, f := opcodeHandlers[opCode]; f {
		return handler(cpu), nil
	}
	errorString := fmt.Sprintf("Unknown opcode %#02x", opCode)
	return 0, errors.New(errorString)
}
